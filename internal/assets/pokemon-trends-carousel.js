import EmblaCarousel from "embla-carousel";
import Autoplay from "embla-carousel-autoplay";

const emblaNode = document.querySelector(".pokedex__trends");
const emblaApi = EmblaCarousel(emblaNode, { loop: true }, [Autoplay()]);

/**
 * @param {import('embla-carousel').EmblaCarouselType} api "EmblaCarousel" instance
 * @returns {() => void} cleanup function
 */
function registerDotEventHandlers(api, emblaRoot) {
  /** @type {HTMLButtonElement[]}  */
  const dots = Array.from(emblaRoot.querySelectorAll(".embla__dot"));

  function registerSlideClickHandler() {
    /** @param {MouseEvent} event */
    function scrollToSlide(event) {
      const dot = event.currentTarget;
      if (!dot) return;

      const index = dot.dataset.index;
      api.scrollTo(+index);
    }

    dots.forEach((dot) => dot.addEventListener("click", scrollToSlide, false));
  }

  function registerSelectedSlideHandler() {
    const prev = api.previousScrollSnap();
    const selected = api.selectedScrollSnap();

    dots[prev].removeAttribute("data-state");
    dots[selected].dataset.state = "selected";
  }

  api
    .on("init", registerSlideClickHandler)
    .on("reInit", registerSlideClickHandler)
    .on("init", registerSelectedSlideHandler)
    .on("reInit", registerSelectedSlideHandler)
    .on("select", registerSelectedSlideHandler);

  return () => {
    dots.forEach((dot) =>
      dot.removeEventListener("click", scrollToSlide, false),
    );
  };
}

const cleanDotsHandlers = registerDotEventHandlers(emblaApi, emblaNode);
emblaApi.on("destroy", cleanDotsHandlers);
