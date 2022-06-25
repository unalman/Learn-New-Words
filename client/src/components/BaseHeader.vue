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
  height: 100%;
  font-size: 1.2rem;
  background-color: rgb(230, 182, 124);
  display: flex;
}
nav ul {
  width: 100%;
  display: flex;
  flex-flow: row wrap;
  box-sizing: border-box;
}
ul li {
  list-style-type: none;
  align-self: center;
  margin-right: auto;
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
  margin-left: 1rem;
}
ul li:nth-child(2) {
  margin-right: 1rem;
}
</style>
