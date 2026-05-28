(() => {
  const themeKey = "theme";
  const menuButton = document.querySelector("[data-menu-toggle]");
  const menuPanel = document.querySelector("[data-menu-panel]");
  const themeButtons = document.querySelectorAll("[data-theme-toggle]");

  const applyTheme = (theme) => {
    const isDark = theme === "dark";
    document.documentElement.classList.toggle("dark", isDark);
    localStorage.setItem(themeKey, theme);
    themeButtons.forEach((button) => {
      button.setAttribute("aria-pressed", String(isDark));
    });
  };

  const initTheme = () => {
    const saved = localStorage.getItem(themeKey);
    const prefersDark = window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches;
    applyTheme(saved || (prefersDark ? "dark" : "light"));
  };

  const setMenuState = (open) => {
    if (!menuPanel || !menuButton) return;
    menuPanel.classList.toggle("is-open", open);
    menuPanel.setAttribute("aria-hidden", String(!open));
    menuButton.setAttribute("aria-expanded", String(open));
    menuButton.setAttribute("aria-label", open ? "Close menu" : "Open menu");
    const icon = menuButton.querySelector("[data-menu-icon]");
    if (icon) {
      icon.classList.toggle("bx-menu", !open);
      icon.classList.toggle("bx-x", open);
    }
  };

  const closeMenuOnDesktop = () => {
    if (window.innerWidth >= 768) {
      setMenuState(false);
    }
  };

  document.addEventListener("DOMContentLoaded", () => {
    initTheme();
    setMenuState(false);

    themeButtons.forEach((button) => {
      button.addEventListener("click", () => {
        const isDark = document.documentElement.classList.contains("dark");
        applyTheme(isDark ? "light" : "dark");
      });
    });

    menuButton?.addEventListener("click", () => {
      const open = !menuPanel?.classList.contains("is-open");
      setMenuState(open);
    });

    document.querySelectorAll("[data-menu-panel] a").forEach((link) => {
      link.addEventListener("click", () => {
        setMenuState(false);
      });
    });
  });

  window.addEventListener("resize", closeMenuOnDesktop);

  document.addEventListener("keydown", (event) => {
    if (event.key === "Escape") {
      setMenuState(false);
    }
  });
})();
