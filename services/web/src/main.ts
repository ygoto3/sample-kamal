async function main() {
  const buttons = document.querySelectorAll('button') as NodeListOf<HTMLButtonElement>;
  buttons.forEach(button => button.addEventListener('click', onClick));
}

async function onClick(ev: MouseEvent) {
  const button = ev.target as null | HTMLButtonElement;
  if (!button) return;
  const id = button.innerText;
  const message = await fetch(`http://api.example.com/v1/users/${id}`).then(res => res.text());
  const messageContainer = document.querySelector('#message') as HTMLDivElement;
  messageContainer.innerHTML = message;
}

main();