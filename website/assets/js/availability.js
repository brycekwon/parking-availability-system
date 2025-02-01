const GREEN = "#22c55e";    // 75%-100% availability
const YELLOW = "#facc15";   // 25%-75% availability
const RED = "#ef4444";      // 0%-25% availability

document.addEventListener("DOMContentLoaded", () => {
    const indicators = document.querySelectorAll("[data-availability]");
    indicators.forEach((indicator) => {
        const ratio = parseFloat(indicator.getAttribute("data-availability"));
        if (ratio >= 0.75) {
            indicator.style.backgroundColor = GREEN;
        } else if (ratio >= 0.25) {
            indicator.style.backgroundColor = YELLOW;
        } else {
            indicator.style.backgroundColor = RED;
        }
    });
});
