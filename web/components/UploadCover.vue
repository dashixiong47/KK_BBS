<template>
  <Dialog v-model="showCover" @confirm="confirm">
    <!-- 选择封面 -->
    <div class="w-[900px] h-96 grid grid-cols-12 gap-5 p-5 overflow-y-auto">
      <div
        v-for="item in coverList"
        class="col-span-2 w-32 h-32 relative cursor-pointer"
      >
        <img
          @click="actived(item)"
          class="w-full h-full object-cover rounded-xl shadow-center hover:border"
          :src="item"
        />
        <div
          v-if="activeCoverList.includes(item)"
          class="w-full h-full pointer-events-none rounded-xl bg-[--bg-masking-color] dark:bg-[--dark-bg-masking-color] absolute top-0 left-0"
        >
          <Icon
            name="ep:check"
            size="2rem"
            class="text-[--illuminate-color] dark:text-[--dark-illuminate-color] absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2"
          />
        </div>
      </div>
      <Upload
        class="w-32 h-32 col-span-2 flex items-center justify-center rounded-xl shadow-center active:shadow-active"
        @uploadSuccess="({ url }) => coverList.push(url)"
      >
        <div>+</div>
      </Upload>
    </div>
  </Dialog>
  <KButton
    class="flex-shrink-0"
    @click="showCover = !showCover"
    :actived="activeCoverList.length"
  >
    选择封面
  </KButton>
</template>
<script setup>

let activeCoverList = ref([]);
const props = defineProps({
  coverList: {
    type: Array,
    default: () => [],
  },
});

let showCover = ref(false);
const emit = defineEmits(["update:modelValue", "selectd"]);
let coverList = ref([...props.coverList]);
const confirm = () => {
  emit("selectd", coverList.value);
};
const actived = (url) => {
  if (activeCoverList.value.includes(url)) {
    activeCoverList.value = activeCoverList.value.filter(
      (item) => item !== url
    );
  } else {
    activeCoverList.value.push(url);
  }
};
</script>
