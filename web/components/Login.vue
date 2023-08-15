<template>
  <div class="bg-[rgba(0,0,0,0.6)] fixed top-0 left-0 w-screen h-screen z-50">
    <div
      class="w-[800px] h-[500px] p-5 bg-white rounded-2xl absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2"
    >
      <p class="w-full flex justify-end">
        <Icon class="cursor-pointer" name="icon-close" @click="close" />
      </p>
      <form>
        <div>
          <input type="text" class="border p-2 rounded-md" />
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { useLoginStore } from "~/stores/main.js";
const store = useLoginStore();
const close = () => {
  store.setLoginStatus();
};
</script>
