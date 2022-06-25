<template>
  <div class="content-container">
    <div class="content-item1">
      Ezberlemek istediğin kelimeleri ekle!
      <div><a href="/addword">Kelime Ekle</a></div>
    </div>
    <div class="content-item2">
      Kendini test et!
      <div><a href="/testword">Kelime Testi</a></div>
    </div>
  </div>
</template>

<style scoped>
.content-container {
  display: grid;
  padding: 2vh;
  grid-gap: 10px;
  box-sizing: border-box;
  height: inherit;
  min-height: inherit;
  align-items: center;
  justify-content: space-around;
}

.content-container > .content-item1,
.content-item2 {
  text-align: center;
  font-size: 2rem;
  box-sizing: border-box;
  width: 40vh;
  border: solid 1px;
  min-width: 180px;
  height: 103px;
}
.content-item1 {
  grid-template-columns: 100px auto;
  min-height: 168px;
}
.content-item2 {
  grid-row-start: 2;
  grid-column-start: 2;
}

div a {
  color: white;
  display: inline-block;
  background-color: #e67e22;
  font-size: 1.7rem;
}

div a:hover {
  color: var(--hotmag);
  text-decoration: none;
}
</style>
