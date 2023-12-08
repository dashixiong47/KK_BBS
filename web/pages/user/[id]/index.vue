<template>
  <div class="">
    <div class="mb-4 flex items-center justify-between">
      <Rolling class="w-[70%]" v-model="selectNode"></Rolling>
      <KSelect
        class="w-[25%]"
        v-model="selectType"
        :options="options"
        @change=""
      ></KSelect>
    </div>
    <List
      class=""
      :type="selectType"
      :userId="route.params.id"
    ></List>
  </div>
</template>

<script setup>
import { useAppConfigStore } from "~/stores/init";
const route = useRoute();
let appConfigStore = useAppConfigStore();
let appConfig = computed(() => appConfigStore.getConfig);
const { typeOptions } = useOptions();
const selectType = ref(-1);
const selectNode = ref(-1);
const options = typeOptions;
useHead({
  title: "动态 - " + appConfig.value.appName,
});
</script>
