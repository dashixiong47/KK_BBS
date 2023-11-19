<template>
  <div class="main">
    <nav class="nav">112</nav>
    <div class="left">
      <Menu />
    </div>
  </div>
</template>

<script setup>
import Menu from "@/components/Menu.vue";
</script>

<style scoped>
.main {
  width: 100vw;
  height: 100vh;
  .nav{
    height: 60px;
    border-bottom: 1px solid #ccc;
  }
  > .left {
    width: 200px;
    height: calc(100vh - 60px);
    background-color: #545c64;
  }
}
</style>
