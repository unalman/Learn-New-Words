<template>
  <div class="addword-container">
    <div class="word-item-1">Kelimeler</div>
    <div class="word-item-2">
      <div class="word-btn">
        <button class="btn-add">Ekle</button>
      </div>
      <div class="word-container">
        <div class="word-header">
          <div class="word-firstLangHeader">Dil 1</div>
          <div class="word-secondLangHeader">Dil 2</div>
        </div>
        <div class="word-list">
          <div class="word-list-item">
            <div class="word-item-firstLang">Yazi1</div>
            <div class="word-item-secondLang">Text1</div>
            <div class="word-item-btns">
              <button class="word-item-btn">Düzenle</button>
              <button class="word-item-btn">Sil</button>
            </div>
          </div>
          <div class="word-list-item">
            <div class="word-item-firstLang">Yazi2</div>
            <div class="word-item-secondLang">Text2</div>
            <div class="word-item-btns">
              <button class="word-item-btn">Düzenle</button>
              <button class="word-item-btn">Sil</button>
            </div>
          </div>
          <div class="word-list-item">
            <div class="word-item-firstLang">Yazi2</div>
            <div class="word-item-secondLang">Text2</div>
            <div class="word-item-btns">
              <button class="word-item-btn">Düzenle</button>
              <button class="word-item-btn">Sil</button>
            </div>
          </div>
          <div class="word-list-item">
            <div class="word-item-firstLang">Yazi2</div>
            <div class="word-item-secondLang">Text2</div>
            <div class="word-item-btns">
              <button class="word-item-btn">Düzenle</button>
              <button class="word-item-btn">Sil</button>
            </div>
          </div>
          <div class="word-list-item">
            <div class="word-item-firstLang">Yazi2</div>
            <div class="word-item-secondLang">Text2</div>
            <div class="word-item-btns">
              <button class="word-item-btn">Düzenle</button>
              <button class="word-item-btn">Sil</button>
            </div>
          </div>
          <div class="word-list-item">
            <div class="word-item-firstLang">Yazi2</div>
            <div class="word-item-secondLang">Text2</div>
            <div class="word-item-btns">
              <button class="word-item-btn">Düzenle</button>
              <button class="word-item-btn">Sil</button>
            </div>
          </div>
          <div class="word-list-item">
            <div class="word-item-firstLang">Yazi2</div>
            <div class="word-item-secondLang">Text2</div>
            <div class="word-item-btns">
              <button class="word-item-btn">Düzenle</button>
              <button class="word-item-btn">Sil</button>
            </div>
          </div>
          <div class="word-list-item">
            <div class="word-item-firstLang">Yazi2</div>
            <div class="word-item-secondLang">Text2</div>
            <div class="word-item-btns">
              <button class="word-item-btn">Düzenle</button>
              <button class="word-item-btn">Sil</button>
            </div>
          </div>
          <div class="word-list-item">
            <div class="word-item-firstLang">Yazi2</div>
            <div class="word-item-secondLang">Text2</div>
            <div class="word-item-btns">
              <button class="word-item-btn">Düzenle</button>
              <button class="word-item-btn">Sil</button>
            </div>
          </div>
          <div class="word-list-item">
            <div class="word-item-firstLang">Yazi2</div>
            <div class="word-item-secondLang">Text2</div>
            <div class="word-item-btns">
              <button class="word-item-btn">Düzenle</button>
              <button class="word-item-btn">Sil</button>
            </div>
          </div>
          <div class="last-item-bottom"></div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
export default {};
</script>
<style scoped>
.addword-container {
  display: grid;
  height: inherit;
  min-height: inherit;
  align-items: center;
  grid-template-columns: 1fr 1fr 8fr 1fr 1fr;
  grid-template-rows: 4rem 1fr;
  box-sizing: border-box;
  overflow: overlay;
}
.addword-container > .word-item-1,
.word-item-2 {
  grid-column: 3;
  box-sizing: border-box;
}
.word-item-1 {
  padding-top: 1.5vh;
  min-height: 38px;
}
.btn-add {
  background-color: var(--dkblue);
  color: white;
  cursor: pointer;
  font-size: 1.4rem;
  border: 1px;
  border-radius: 1rem;
  height: 2.5rem;
}
.word-btn {
  padding-top: 1.5rem;
  padding-right: 4rem;
  text-align: end;
}
.word-container {
  margin-top: 1rem;
  font-size: 1.4rem;
}
.word-header {
  display: flex;
  align-items: center;
  padding-bottom: 0.8rem;
  justify-content: space-between;
}
.word-firstLangHeader,
.work-secondLangHeader {
  width: 4rem;
}
.word-firstLangHeader {
  padding-left: 10rem;
  padding-right: 14rem;
  
}
.word-secondLangHeader {
  padding-right: 39rem;
  width: 4rem;
}
.word-list-item {
  display: flex;
  align-items: center;
  padding-top: 0.7rem;
  padding-bottom: 0.7rem;
  border-top: 1px solid #9eb5cc;
  justify-content: space-between;
}
.last-item-bottom {
  border-top: 1px solid #9eb5cc;
}
.word-item-firstLang {
  padding-left: 10rem;
  padding-right: 1rem;
}
.word-item-secondLang {
  padding-right: 10rem;
}
.word-item-btns {
  padding-right: 5rem;
  padding-left: 0.7rem;
}
.word-item-btn {
  background-color: var(--dkblue);
  color: white;
  cursor: pointer;
  font-size: 1.4rem;
  border: 1px;
  border-radius: 1rem;
  margin: 6px;
  box-sizing: border-box;
}
</style>
