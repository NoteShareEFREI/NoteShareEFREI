

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
        content.innerHTML = '&lt;&lt; Study sheet &gt;&gt;';
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

const sheetComments = {
    'Study sheet 1': [
        { user: 'User1', text: 'The study sheet is good but there is a point that need to be reworked.' }
    ],
    'Study sheet 2': [],
    'Study sheet 3': [
        { user: 'User2', text: 'Great sheet, very clear!' }
    ]
};

const sheetData = {
    'Study sheet 1': { pdf: 'pdfs/bulletin-may26.pdf' },
    'Study sheet 2': { pdf: 'pdfs/Bee_Movie.pdf' },
    'Study sheet 3': { pdf: 'pdfs/The_Chapel_on_the_Cliffs_v1.0.pdf' },
};