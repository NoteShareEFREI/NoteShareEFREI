function renderSidebar() {
    const sidebar = document.querySelector('.folders');
    sidebar.innerHTML = '';

    Object.entries(folderData).forEach(([folderName, sheets], index) => {
        const isFirst = index === 0;
        const folderId = 'folder' + (index + 1);
        const folderEl = document.createElement('div');
        folderEl.className = 'folder-item' + (isFirst ? '' : ' folder-collapsed');
        folderEl.id = folderId;

        folderEl.innerHTML = `
            <div class="folder-header" onclick="toggleFolder('${folderId}')">
                <span class="folder-toggle">${isFirst ? '−' : '+'}</span>
                <span>${folderName}</span>
            </div>
            <div class="folder-children">
                ${sheets.map(sheet => `
                    <span class="sheet-link" onclick="loadSheet('${sheet}')">${sheet}</span>
                `).join('')}
            </div>
        `;

        sidebar.appendChild(folderEl);
    });
}

function toggleFolder(id) {
    const el = document.getElementById(id);
    const isCollapsed = el.classList.contains('folder-collapsed');
    const toggle = el.querySelector('.folder-toggle');
    if (isCollapsed) {
        el.classList.remove('folder-collapsed');
        toggle.textContent = '−';
    } 
    else {
        el.classList.add('folder-collapsed');
        toggle.textContent = '+';
    }
}

function createCommentCard(user, text) {
  const card = document.createElement('div');
  card.className = 'comment-card';
  card.innerHTML = `
    <div class="comment-top">
      <span class="comment-username">${user}</span>
    </div>
    <p class="comment-text">${text}</p>
  `;
  return card;
}

function renderComments(sheetName) {
  const wrap = document.querySelector('.comment-input-wrap');

  // Remove all existing comment cards
  document.querySelectorAll('.comment-card').forEach(card => card.remove());

  // Re-render cards from the data store
  const comments = sheetComments[sheetName] || [];
  comments.forEach(comment => {
    const card = createCommentCard(comment.user, comment.text);
    wrap.after(card);
  });
}

let currentSheet = null;

function loadSheet(name) {
    currentSheet = name;
    document.getElementById('sheet-title').textContent = name;
    
    const content = document.getElementById('sheet-display');
    const sheet = sheetData[name];

    if (sheet && sheet.pdf) {
        content.innerHTML = `
        <iframe
            src="${sheet.pdf}"
            width="100%"
            height="100%"
            style="border: none; border-radius: 12px;"
        ></iframe>
        `;
    } else {
        content.innerHTML = '&lt;&lt;ERROR : PDF of study sheet not found !&gt;&gt;';
    }
    renderComments(name)
}

let commentsCollapsed = false;
function toggleComments() {
    const panel = document.getElementById('comments-panel');
    commentsCollapsed = !commentsCollapsed;
    if (commentsCollapsed) {
        panel.classList.add('comments-collapsed');
    } 
    else {
        panel.classList.remove('comments-collapsed');
    }
}

let addCommentVisible = false;
function displayAddComments() {
    const input_comment = document.getElementById('adding-comment');
    addCommentVisible = !addCommentVisible;
    input_comment.style.display = addCommentVisible ? 'flex' : 'none';
}

function handleCommentInput(event) {
    if (event.key !== 'Enter') return;

    const input = event.target;
    const text = input.value.trim();
    if (!text) return;

    if (!currentSheet) return;

    if (!sheetComments[currentSheet]) sheetComments[currentSheet] = [];
    sheetComments[currentSheet].push({ user: 'You', text });

    const card = createCommentCard('You', text);
    const wrap = document.querySelector('.comment-input-wrap');
    wrap.after(card);

    input.value = '';
    addCommentVisible = false;
    document.getElementById('adding-comment').style.display = 'none';
}

function handleSearch(event) {
    const query = event.target.value.trim().toLowerCase();

    document.querySelectorAll('.folder-item').forEach(folder => {
        const folderName = folder.querySelector('.folder-header span:last-child').textContent.toLowerCase();
        const sheets = folder.querySelectorAll('.sheet-link');

        let anySheetMatches = false;

        sheets.forEach(sheet => {
            const sheetName = sheet.textContent.toLowerCase();
            const matches = sheetName.includes(query) || folderName.includes(query);
            sheet.style.display = matches ? 'block' : 'none';
            if (matches) anySheetMatches = true;
        });

        // Show the folder if its name matches or any of its sheets match
        const folderMatches = folderName.includes(query) || anySheetMatches;
        folder.style.display = folderMatches ? 'block' : 'none';

        // Auto-expand folder if a sheet inside it matches
        if (anySheetMatches && query !== '') {
            folder.classList.remove('folder-collapsed');
            folder.querySelector('.folder-toggle').textContent = '−';
        }
    });

    // If search is cleared, restore everything
    if (query === '') restoreSidebar();
}

function restoreSidebar() {
    document.querySelectorAll('.folder-item').forEach(folder => {
        folder.style.display = 'block';
        folder.querySelectorAll('.sheet-link').forEach(sheet => {
            sheet.style.display = 'block';
        });
    });
}

const sheetComments = {
    'Study sheet 1': [
        { user: 'User1', text: 'The study sheet is good but there is a point that need to be reworked.' }
    ],
    'Study sheet 2': [],
    'Study sheet 3': [
        { user: 'User2', text: 'Great sheet, very clear!' }
    ]
};

const folderData = {
    'Folder1': ['Study sheet 1', 'Study sheet 2', 'Study sheet 3'],
    'Folder2': ['Sheet A'],
    'Folder3': ['Sheet B'],
};

document.addEventListener('DOMContentLoaded', () => {
    renderSidebar();
    checkAdminStatus();
});

// Check if user is admin and show/hide admin button
function checkAdminStatus() {
    // Check if admin button exists
    const adminBtn = document.getElementById('admin-btn');
    if (!adminBtn) return;

    // Try to fetch admin status from an API endpoint or check from localStorage
    // Since we need to check the user's admin status, we can make a request
    fetch('/api/check-admin', {
        method: 'GET',
        credentials: 'include'
    })
    .then(response => {
        if (response.ok) {
            return response.json();
        }
        return { isAdmin: false };
    })
    .then(data => {
        if (data.isAdmin) {
            adminBtn.style.display = 'block';
        }
    })
    .catch(error => {
        console.log('Could not check admin status:', error);
        adminBtn.style.display = 'none';
    });
}
