(() => {
  const themeKey = "theme";
  const menuButton = document.querySelector("[data-menu-toggle]");
  const menuPanel = document.querySelector("[data-menu-panel]");
  const themeButtons = document.querySelectorAll("[data-theme-toggle]");
  const copyButtons = document.querySelectorAll("[data-code-copy]");

  const setDocumentTheme = (isDark) => {
    document.documentElement.classList.toggle("dark", isDark);
    document.documentElement.style.colorScheme = isDark ? "dark" : "light";
  };

  const applyTheme = (theme) => {
    const isDark = theme === "dark";
    setDocumentTheme(isDark);
    localStorage.setItem(themeKey, theme);
    themeButtons.forEach((button) => {
      button.setAttribute("aria-pressed", String(isDark));
    });
  };

  const applyThemeWithoutCTAAnimation = (theme) => {
    document.documentElement.classList.add("is-theme-switching");
    applyTheme(theme);
    window.requestAnimationFrame(() => {
      window.requestAnimationFrame(() => {
        document.documentElement.classList.remove("is-theme-switching");
      });
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
      if (!menuPanel || !menuButton) return;
      menuPanel.classList.remove("is-open");
      menuPanel.setAttribute("aria-hidden", "false");
      menuButton.setAttribute("aria-expanded", "false");
      menuButton.setAttribute("aria-label", "Open menu");
    }
  };

  const copyToClipboard = async (text) => {
    if (navigator.clipboard?.writeText) {
      await navigator.clipboard.writeText(text);
      return;
    }

    const textarea = document.createElement("textarea");
    textarea.value = text;
    textarea.setAttribute("readonly", "");
    textarea.style.position = "absolute";
    textarea.style.left = "-9999px";
    document.body.appendChild(textarea);
    textarea.select();
    document.execCommand("copy");
    document.body.removeChild(textarea);
  };

  const handleCodeCopy = async (button) => {
    const block = button.closest("[data-code-block]");
    const code = block?.querySelector("code");
    const codeLines = code ? Array.from(code.querySelectorAll(".cl")) : [];
    const text = (codeLines.length > 0
      ? codeLines.map((line) => line.innerText.replace(/\u200b/g, "")).join("\n")
      : code?.innerText
    )?.replace(/\n$/, "");

    if (!text) return;

    try {
      await copyToClipboard(text);
      button.textContent = "Copied";
      button.dataset.copyState = "copied";
      window.setTimeout(() => {
        button.textContent = "Copy";
        button.dataset.copyState = "idle";
      }, 1800);
    } catch {
      button.textContent = "Failed";
      button.dataset.copyState = "idle";
    }
  };

  document.addEventListener("DOMContentLoaded", () => {
    initTheme();
    if (window.innerWidth < 768) {
      setMenuState(false);
    } else {
      closeMenuOnDesktop();
    }

    themeButtons.forEach((button) => {
      button.addEventListener("click", () => {
        const isDark = document.documentElement.classList.contains("dark");
        applyThemeWithoutCTAAnimation(isDark ? "light" : "dark");
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

    copyButtons.forEach((button) => {
      button.addEventListener("click", () => {
        handleCodeCopy(button);
      });
    });
  });

  window.addEventListener("resize", closeMenuOnDesktop);

  document.addEventListener("click", (event) => {
    if (!menuPanel || !menuButton) return;
    if (window.innerWidth >= 768) return;
    if (!menuPanel.classList.contains("is-open")) return;
    if (menuPanel.contains(event.target) || menuButton.contains(event.target)) return;
    setMenuState(false);
  });

  document.addEventListener("keydown", (event) => {
    if (event.key === "Escape") {
      setMenuState(false);
    }
  });
})();
