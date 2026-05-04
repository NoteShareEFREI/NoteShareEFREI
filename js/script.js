function createCommentCard(pseudo, text) {
  const card = document.createElement('div');
  card.className = 'comment-card';
  card.innerHTML = `
    <div class="comment-top">
      <span class="comment-username">${pseudo}</span>
    </div>
    <p class="comment-text">${text}</p>
  `;
  return card;
}

function renderComments(sheetName) {
  const container = document.getElementById('comments-container');
  container.innerHTML = ''; // Clear existing

  fetch(`/api/comments?sheet=${sheetName}`)
    .then(response => response.json())
    .then(comments => {
      comments.forEach(comment => {
        const card = createCommentCard(comment.Pseudo, comment.Content);
        container.appendChild(card);
      });
    })
    .catch(error => console.error('Error loading comments:', error));
}

let currentSheet = null;

function handleCommentInput(event) {
    if (event.key !== 'Enter') return;

    const input = event.target;
    const text = input.value.trim();
    if (!text) return;

    if (!currentSheet) return;

    // Post comment
    const formData = new FormData();
    formData.append('sheet', currentSheet);
    formData.append('content', text);

    fetch('/api/comments/add', {
        method: 'POST',
        body: formData,
        credentials: 'include'
    })
    .then(response => {
        if (response.ok) {
            // Re-render comments
            renderComments(currentSheet);
            input.value = '';
        } else {
            alert('Failed to add comment');
        }
    })
    .catch(error => console.error('Error adding comment:', error));
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

document.addEventListener('DOMContentLoaded', () => {
    checkAdminStatus();
});
