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

const autoRunHtmx = () => {
  const form = document.getElementById('cardform');
  const nrleft = +form.times.value;

  if (nrleft > 0) {
    const data = {
      name: `John ${nrleft} Htmx`,
      email: `john.${nrleft}@htmx.com`,
      times: nrleft - 1,
    };

    fillForm(form, data);
    const button = document.getElementById('card-button-htmx');
    button.click();

    document.addEventListener('htmx:afterSwap', () => {
      autoRunHtmx();
    }, { once: true });
  } else {
    const endTime = +new Date();
    const duration = endTime - STATE.startTime;
    const durationEl = document.getElementById('autorun-result');
    durationEl.innerHTML = `${duration} ms`;
  }
};

const autoRunJson = () => {
  const form = document.getElementById('cardform');
  const nrleft = +form.times.value;

  if (nrleft > 0) {
    const data = {
      name: `John ${nrleft} Json`,
      email: `john.${nrleft}@json.com`,
      times: nrleft - 1,
    };

    fillForm(form, data);
    const button = document.getElementById('card-button-json');
    button.click();

    document.addEventListener('json:ready', () => {
      autoRunJson();
    }, { once: true });
  } else {
    const endTime = +new Date();
    const duration = endTime - STATE.startTime;
    const durationEl = document.getElementById('autorun-result');
    durationEl.innerHTML = `${duration} ms`;
  }
};

const autoStartHtmx = () => {
  STATE.startTime = +new Date();
  autoRunHtmx();
};

const autoStartJson = () => {
  STATE.startTime = +new Date();
  autoRunJson();
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

      const event = new CustomEvent('json:ready', { bubbles: true });
      document.dispatchEvent(event);
    });
  } else if (e.target.matches('#autorun-htmx, #autorun-htmx *')) {
    autoStartHtmx();
  } else if (e.target.matches('#autorun-json, #autorun-json *')) {
    autoStartJson();
  }
});

