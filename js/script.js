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

function loadSheet(name) {
    document.getElementById('sheet-display').textContent = '<< ' + name + ' >>';
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

let add_comments = true;
function displayAddComments() {
    const input_commnent = document.getElementById('adding-comment');
    commentsCollapsed = !commentsCollapsed;
    if (commentsCollapsed) {
        input_commnent.style.display = 'flex'
    } 
    else {
        input_commnent.style.display = 'none'
    }
}

function handleCommentInput(event) {
    if (event.key !== 'Enter') return;

    const input = event.target;
    const text = input.value.trim();
    if (!text) return;

    const card = document.createElement('div');
    card.className = 'comment-card';
    card.innerHTML = `
    <div class="comment-top">
        <span class="comment-username">You</span>
    </div>
    <p class="comment-text">${text}</p>
    `;

    const wrap = document.querySelector('.comment-input-wrap');
    wrap.after(card);

    input.value = '';
    displayAddComments();
}