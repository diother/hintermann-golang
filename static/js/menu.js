const content = document.querySelector("[data-content]");
const mobileMenu = document.querySelector("[data-menu]");
const openBtn = document.querySelector('[data-menu-toggle="open"]');
const closeBtn = document.querySelector('[data-menu-toggle="close"]');
let scrollTop = 0;

openBtn.addEventListener("click", () => {
    scrollTop = window.scrollY;
    content.style.position = "fixed";
    content.style.overflow = "hidden";
    mobileMenu.classList.remove("hidden");
    openBtn.classList.add("hidden");
    closeBtn.classList.remove("hidden");
});

closeBtn.addEventListener("click", () => {
    content.removeAttribute("style");
    mobileMenu.classList.add("hidden");
    openBtn.classList.remove("hidden");
    closeBtn.classList.add("hidden");
    window.scrollTo(0, scrollTop);
});
