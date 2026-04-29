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