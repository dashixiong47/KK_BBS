<template>
  <!-- 导航 -->
  <PcNav />
  <div
    class="h-full grid grid-cols-24 p-3 sm:p-5 overflow-y-scroll relative z-0"
  >
    <LeftBox class="fixed left-5 hidden xl:block" />
    <main
      class="mx-auto col-start-1 col-span-24 xl:col-start-6 xl:col-end-20 grid grid-cols-3"
    >
      <div class="h-full sm:mr-2 col-span-3 sm:col-span-2">
        <slot />
      </div>
      <AdditionalPanel class="hidden sm:block col-span-1 h-full" />
    </main>
    <RightBox class="fixed right-5 hidden xl:block" />
  </div>
</template>

<script setup>
import { useAppConfigStore } from "~/stores/init";
let appConfigStore = useAppConfigStore();
let appConfig = computed(() => appConfigStore.getConfig);
useHead({
  title: appConfig.value.appName,
});
</script>
