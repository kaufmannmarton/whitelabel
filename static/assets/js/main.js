window.onload = function () {
    var hamburgerIconEl = document.querySelector("#hamburger-icon");

    hamburgerIconEl.addEventListener("click", function() {
        this.classList.toggle("active");
    });

    // Scroll to section when a nav link/button is pressed
    document.querySelectorAll("[data-container]").forEach(function(element) {
        element.addEventListener("click", function() {
            document.querySelector("#" + this.dataset.container).scrollIntoView();

            // Close navigation - applies on small screens only
            if (window.getComputedStyle(hamburgerIconEl).display !== "none") {
                hamburgerIconEl.classList.toggle("active");
            }
        });
    });

    // Load twitter widget when photos get into view
    var photoContainerEl = document.querySelector("#photo-container");

    var observer = new IntersectionObserver(function (entries, observer) {
        entries.forEach(function (entry) {
            if (entry.target === photoContainerEl && entry.isIntersecting) {
                observer.unobserve(photoContainerEl);
                document.querySelector("#twitter-timeline").classList.add("twitter-timeline");

                twttr.widgets.load();
            }
        })
    }, {
            root: null,
            rootMargin: "50px 0px 0px 0px",
            threshold: 0
        }
    );

    observer.observe(photoContainerEl);
}
