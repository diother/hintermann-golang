document.addEventListener("DOMContentLoaded", () => {
    const content = document.querySelector(".content");
    const mobileMenu = document.querySelector(".mobile-menu");
    const openBtn = document.querySelector(".mobile-menu__open");
    const closeBtn = document.querySelector(".mobile-menu__close");
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
});
