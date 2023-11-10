<template>
  <!-- 导航 -->
  <PcNav />
  <div class="overflow-y-auto relative flex justify-between p-2 xl:p-5">
    <LeftBox class="sticky top-0 left-0 w-[20%] hidden xl:block" />
    <main class="w-full flex xl:w-[60%]">
      <div class="w-full sm:w-[70%] sm:pr-5 md:w-[65%]">
        <slot />
      </div>
      <AdditionalPanel class="w-[30%] hidden sm:block md:w-[35%]" />
    </main>
    <RightBox class="sticky top-0 right-0 w-[20%] hidden xl:flex" />
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
