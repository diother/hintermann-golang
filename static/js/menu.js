document.addEventListener("DOMContentLoaded", () => {
    const wrapper = document.querySelector("#content-wrapper");
    const openBtn = document.querySelector("#open");
    const closeBtn = document.querySelector("#close");
    const mobileMenu = document.querySelector("#mobile-menu");
    let scrollTop = 0;

    openBtn.addEventListener("click", () => {
        scrollTop = window.scrollY;
        wrapper.classList.add("fixed", "overflow-auto");
        openBtn.classList.add("hidden");
        closeBtn.classList.remove("hidden");
        mobileMenu.classList.remove("hidden");
        wrapper.scrollTo(0, scrollTop);
    });

    closeBtn.addEventListener("click", () => {
        wrapper.classList.remove("fixed", "overflow-auto");
        openBtn.classList.remove("hidden");
        closeBtn.classList.add("hidden");
        mobileMenu.classList.add("hidden");
        window.scrollTo(0, scrollTop);
    });
});
