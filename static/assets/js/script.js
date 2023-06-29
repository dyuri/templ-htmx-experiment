const cardWidget = (card) => {
  return `
    <dy-card>
      <div class="name">${card.name}</div>
      <div class="email">${card.email}</div
    </dy-card>
  `;
};

const fillForm = (form, data) => {
  const keys = Object.keys(data);
  keys.forEach((key) => {
    const input = form.querySelector(`[name="${key}"]`);
    if (input) {
      input.value = data[key];
    }
  });
};

const STATE = {
  startTime: 0
};

const autoRun = (type) => {
  const form = document.getElementById('cardform');
  const nrleft = +form.times.value;

  if (nrleft > 0) {
    const data = {
      name: `John ${nrleft} ${type.toUpperCase()}`,
      email: `john.${nrleft}@${type}.com`,
      times: nrleft - 1,
    };

    fillForm(form, data);
    const button = document.getElementById(`card-button-${type}`);
    button.click();

    if (type === 'htmx') {
      document.addEventListener('htmx:afterSwap', () => {
        autoRun(type);
      }, { once: true });
    } else {
      document.addEventListener('dypost:ready', () => {
        autoRun(type);
      }, { once: true });
    }
  } else {
    const endTime = +new Date();
    const duration = endTime - STATE.startTime;
    const durationEl = document.getElementById('autorun-result');
    durationEl.innerHTML = `${duration} ms`;
    form.times.value = STATE.times;
  }
};

const autoStart = (type) => {
  const form = document.getElementById('cardform');

  STATE.startTime = +new Date();
  STATE.times = +form.times.value;

  autoRun(type);
};

document.addEventListener('click', (e) => {
  if (e.target.matches('[dy-post], [dy-post] *')) {
    e.preventDefault();
    const button = e.target.closest('[dy-post]');
    const id = button.getAttribute('dy-post');
    const url = button.getAttribute('dy-post-target');
    const type = button.getAttribute('dy-post-type');
    const form = document.getElementById(id);
    const formData = new FormData(form);

    fetch(url, {
      method: 'POST',
      body: formData,
    }).then((response) => {
      if (response.ok) {
        return type === 'json' ? response.json() : response.text();
      }
      throw new Error('Network response was not ok.');
    }).then((data) => {
      const cardCnt = document.getElementById('card_cnt');
      cardCnt.innerHTML = type === 'json' ? cardWidget(data) : data;

      const event = new CustomEvent('dypost:ready', { bubbles: true });
      document.dispatchEvent(event);
    });
  } else if (e.target.matches('[autorun]')) {
    autoStart(e.target.getAttribute('autorun'));
  }
});

