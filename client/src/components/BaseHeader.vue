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
  font-size: 1.2rem;
  background-color: rgb(230, 182, 124);
  display: flex;
}
nav ul {
  width: 100%;
  display: flex;
  flex-flow: row wrap;
}
ul li {
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
  margin-right: auto;
}
ul li:nth-child(2) {
  margin-right: 2rem;
}
</style>
