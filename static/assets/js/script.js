const cardWidget = (card) => {
  return `
    <dy-card>
      <div class="name">${card.name}</div>
      <div class="email">${card.email}</div
    </dy-card>
  `;
};

document.addEventListener('click', (e) => {
  if (e.target.matches('[dy-post], [dy-post] *')) {
    e.preventDefault();
    const button = e.target.closest('[dy-post]');
    const id = button.getAttribute('dy-post');
    const url = button.getAttribute('dy-post-target');
    const form = document.getElementById(id);
    const formData = new FormData(form);

    fetch(url, {
      method: 'POST',
      body: formData,
    }).then((response) => {
      if (response.ok) {
        return response.json();
      }
      throw new Error('Network response was not ok.');
    }).then((data) => {
      const cardCnt = document.getElementById('card_cnt');
      cardCnt.innerHTML = cardWidget(data);
    });
  }
});

document.addEventListener('htmx:afterSettle', () => {
  console.log('htmx:afterSettle');
});
