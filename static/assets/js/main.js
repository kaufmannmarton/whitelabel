window.onload = function () {
    var photoContainerEl = document.querySelector("#photo-container");

    var observer = new IntersectionObserver(function (entries, observer) {
        entries.forEach(function (entry) {
            if (entry.target === photoContainerEl && entry.isIntersecting) {
                observer.unobserve(photoContainerEl);
                document.querySelector("#twitter-timeline").classList.add('twitter-timeline');

                twttr.widgets.load();
            }
        })
    }, {
            root: null,
            rootMargin: '50px 0px 0px 0px',
            threshold: 0
        }
    );

    observer.observe(photoContainerEl);
}
