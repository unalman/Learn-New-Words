<template>
  <div class="content-container">
    <div class="content-item1">
      <div>Ezberlemek istediğin kelimeleri ekle!</div>
      <div class="content-btnContainer">
        <a class="content-wordbtn" href="/addword">Kelime Ekle</a>
      </div>
    </div>
    <div class="content-item2">
      <div>Kendini test et!</div>
      <div class="content-btnContainer">
        <a class="content-wordbtn" href="/wordtest">Kelime Testi</a>
      </div>
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
  font-size: 3.5vh;
  box-sizing: border-box;
  width: 40vh;
  border: solid 0.4px;
  border-color: gray;
  min-width: 180px;
  height: 14vh;
  border-radius: 1rem;
  padding-top: 1rem;
  background-color: antiquewhite;
}
.content-item1 {
  grid-template-columns: 100px auto;
  min-height: 144px;
}
.content-item2 {
  grid-row-start: 2;
  grid-column-start: 2;
  min-height: 101px;
}
.content-btnContainer {
  padding-top: 1rem;
  text-align: center;
}
.content-wordbtn {
  color: white;
  display: inline-block;
  background-color: #e67e22;
  font-size: 1.7rem;
  border-radius: 0.7rem;
  width: 10rem;
}
div .content-wordbtn:hover {
  color: var(--hotmag);
  text-decoration: none;
}
</style>
