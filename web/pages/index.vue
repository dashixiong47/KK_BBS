<script setup>
const { $imageViewer } = useNuxtApp();
import { useRouter } from "nuxt/app";
import useMobileDetect from "~/composables/useMobileDetect";
let router = useRouter();
const { locale } = useI18n();

const { isMobile } = useMobileDetect();
const runtimeConfig = useRuntimeConfig();
function test() {
  console.log("test");
}
function openViewer(index = 0) {
  $imageViewer.open(
    [
      {
        url: "http://localhost:8080/files/images/2023/11/24/81ca514d094574c50754b75b8ae4f790.png",
      },
    ],
    index
  );
}
</script>
<template>
  <div class="home grid grid-cols-12 gap-5">
    <div class="col-span-4 shadow-center rounded-xl overflow-hidden">
      <img @click="openViewer(0)" src="/images/bg.png" alt="" srcset="" />
    </div>
    <div class="col-span-8 relative shadow-center rounded-xl">
      热点
      <div class="absolute w-full h-20 -bottom-10 bg-[--bg-color]"></div>
    </div>
    <div class="border h-96 col-span-12 shadow-center rounded-xl">测试</div>
  </div>
</template>
<style lang="postcss"></style>
