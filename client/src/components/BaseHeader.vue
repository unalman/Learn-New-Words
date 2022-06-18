<template>
  <nav>
    <ul>
      <li><a href="/">Anasayfa</a></li>
      <li><a href="/contact">İletişim</a></li>
    </ul>
  </nav>
</template>

<style scoped>
nav {
  font-size: 2.1vh;
  background-color: rgb(230, 182, 124);
  display: flex;
}
nav ul {
  height: 100%;
  width: 100%;
  display: flex;
  flex-flow: row wrap;
  justify-content: center;
  align-items: center;
  gap: 2rem;
  box-sizing: border-box;
}
ul li {
  display: block;
  list-style-type: none;
}
nav a {
  text-decoration: none;
  color: var(--white);
}
nav a:hover {
  color: var(--hotmag);
  text-decoration: none;
}
nav ul li:first-child {
  flex-basis: auto;
  margin-right: auto;
  order: 0;
}
ul li:nth-child(2) {
  order: 2;
  margin-right: 1rem;
}
</style>
