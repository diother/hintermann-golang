document.addEventListener("DOMContentLoaded", () => {
    const wrapper = document.querySelector("#content");
    const openBtn = document.querySelector("#open");
    const closeBtn = document.querySelector("#close");
    const mobileMenu = document.querySelector("#mobile-menu");
    let scrollTop = 0;

    openBtn.addEventListener("click", () => {
        scrollTop = window.scrollY;
        wrapper.classList.add("fixed", "overflow-hidden");
        openBtn.classList.add("hidden");
        closeBtn.classList.remove("hidden");
        mobileMenu.classList.remove("hidden");
    });

    closeBtn.addEventListener("click", () => {
        wrapper.classList.remove("fixed", "overflow-hidden");
        openBtn.classList.remove("hidden");
        closeBtn.classList.add("hidden");
        mobileMenu.classList.add("hidden");
        window.scrollTo(0, scrollTop);
    });
});
