const keyAuth = "telin-auth";
const loader = `
<div class="d-flex justify-content-center">
  <div class="spinner-border" role="status">
    <span class="sr-only">Loading...</span>
  </div>
</div>
`;

function cekLogedIn() {
  var user = localStorage.getItem(keyAuth);
  if (user) {
    var obj = JSON.parse(user);
    if (obj !== null) {
      return obj;
    }
  } else {
    return false;
  }
}
function logOut() {
  localStorage.removeItem(keyAuth);
}

function demo() {}

const images = document.querySelectorAll("img[data-src]");
const options = {
  root: null,
  rootMargin: "0px",
  threshold: 0.1,
};

const handleIntersection = (entries, observer) => {
  entries.forEach((entry) => {
    if (entry.isIntersecting) {
      const image = entry.target;
      const src = image.getAttribute("data-src");
      image.src = src;

      observer.unobserve(image);
    }
  });
};

const observer = new IntersectionObserver(handleIntersection, options);
images.forEach((image) => {
  observer.observe(image);
});

function Numbering(string) {
  return new Intl.NumberFormat().format(string);
}

function SaveLocalStorage(key, object) {
  localStorage.setItem(key, JSON.stringify(object));
}

function GetLocalStorage(key) {
  const objectString = localStorage.getItem(key);
  return JSON.parse(objectString);
}
