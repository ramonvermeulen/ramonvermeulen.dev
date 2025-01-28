function updateThemeIcon(isDarkMode) {
    const iconElement = document.querySelector('.bx');
    if (iconElement) {
        if (isDarkMode) {
            iconElement.classList.replace('bxs-moon', 'bxs-sun');
        } else {
            iconElement.classList.replace('bxs-sun', 'bxs-moon');
        }
    }
}

function applyTheme() {
    const isDarkMode = localStorage.currentTheme === "dark" ||
        (!("theme" in localStorage) && window.matchMedia("(prefers-color-scheme: dark)").matches);
    document.documentElement.classList.toggle("dark", isDarkMode);
    updateThemeIcon(isDarkMode);
}

function toggleTheme() {
    const isDarkMode = document.documentElement.classList.contains("dark");
    if (isDarkMode) {
        localStorage.currentTheme = "light";
    } else {
        localStorage.currentTheme = "dark";
    }
    applyTheme();
}

document.addEventListener("DOMContentLoaded", applyTheme);
