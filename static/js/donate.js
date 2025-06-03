const toggleButtons = document.querySelectorAll('[data-donate-toggle]');
const onceButton = document.querySelector('[data-donate-toggle="once"]');
const monthlyButton = document.querySelector('[data-donate-toggle="monthly"]');
const formOnce = document.getElementById('form-once');
const formMonthly = document.getElementById('form-monthly');

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
    "pay-1": "https://buy.stripe.com/bJeeVdbSW34Bbfn8Lt93y01",
    "pay-2": "https://buy.stripe.com/5kQ5kD1ei7kRerz2n593y02",
    "pay-3": "https://buy.stripe.com/cNibJ16yCeNjbfn6Dl93y03",
    "pay-4": "https://buy.stripe.com/eVq3cv0ae48F1ENf9R93y04",
    "pay-5": "https://buy.stripe.com/6oU7sL2im34BgzH6Dl93y05",
    "sub-1": "https://buy.stripe.com/3cI28r6yCdJfdnv7Hp93y06",
    "sub-2": "https://buy.stripe.com/00wdR99KOfRn5V36Dl93y07",
    "sub-3": "https://buy.stripe.com/aFaeVd4queNj6Z78Lt93y08",
    "sub-4": "https://buy.stripe.com/bJe28r0ae20x4QZ2n593y0a",
    "sub-5": "https://buy.stripe.com/14AcN5g9cdJf4QZ6Dl93y09"
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
