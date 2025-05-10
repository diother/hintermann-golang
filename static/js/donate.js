const toggleButtons = document.querySelectorAll('[data-donate-toggle]');
const onceButton = document.querySelector('[data-donate-toggle="once"]');
const monthlyButton = document.querySelector('[data-donate-toggle="monthly"]');
const formOnce = document.getElementById('form-once');
const formMonthly = document.getElementById('form-monthly');

console.log(onceButton)

toggleButtons.forEach(button => {
    button.addEventListener('click', () => {
        const target = button.dataset.donateToggle;

        if (target === 'once') {
            formOnce.classList.remove('hidden');
            onceButton.classList.add('donate__toggle--active');
            formMonthly.classList.add('hidden');
            monthlyButton.classList.remove('donate__toggle--active');
        } else if (target === 'monthly') {
            formOnce.classList.add('hidden');
            onceButton.classList.remove('donate__toggle--active');
            formMonthly.classList.remove('hidden');
            monthlyButton.classList.add('donate__toggle--active');
        }
    });
});

const links = {
    "pay-1": "/donate/thank-you",
    "pay-2": "/donate/thank-you",
    "pay-3": "/donate/thank-you",
    "pay-4": "/donate/thank-you",
    "pay-5": "/donate/thank-you",
    "sub-1": "/donate/thank-you",
    "sub-2": "/donate/thank-you",
    "sub-3": "/donate/thank-you",
    "sub-4": "/donate/thank-you",
    "sub-5": "/donate/thank-you"
};

const handleFormSubmit = (form) => {
    form.addEventListener('submit', (e) => {
        e.preventDefault();

        const selectedOption = form.querySelector('input[name="option"]:checked');
        const selectedValue = selectedOption.value;
        const redirectUrl = links[selectedValue];

        if (redirectUrl) {
            window.location.href = redirectUrl;
        }
    });
};

handleFormSubmit(formOnce);
handleFormSubmit(formMonthly);
