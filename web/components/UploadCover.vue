<template>
  <!-- 选择封面 -->
  <div class="w-[800px] h-96 grid grid-cols-12 gap-2 px-5 m-5 overflow-y-auto">
    <div
      v-for="item in [...modelValue]"
      class="col-span-2 w-32 h-32 relative cursor-pointer"
    >
      <img
        @click="actived(item)"
        class="w-full h-full object-cover rounded-xl shadow-center hover:border"
        :src="item"
      />
      <div
        v-if="activeCoverList.includes(item)"
        class="w-full h-full pointer-events-none bg-[--masking-color] dark:bg-[--dark-masking-color] absolute top-0 left-0"
      >
        <Icon
          name="ep:check"
          size="2rem"
          class="text-[--color-white] dark:text-[--dark-color-white] absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2"
        />
      </div>
    </div>
    <div
      class="w-32 h-32 col-span-2 flex items-center justify-center rounded-xl shadow-center active:shadow-active"
    >
      <Upload
        class="flex items-center justify-center"
        @uploadSuccess="({ url }) => modelValue.push(url)"
      >
        +
      </Upload>
    </div>
  </div>
  <div class="flex justify-end">
    <KButton @click="selected" class="primary-text"> 确定 </KButton>
  </div>
</template>
<script setup>
let activeCoverList = ref([]);
const props = defineProps({
  modelValue: {
    type: Array,
    default: ()=>[],
  },
});
const emit = defineEmits(["update:modelValue","selectd"]);
const coverList = ref(props.coverList);
const modelValue=computed({
    get:()=>props.modelValue,
    set:(value)=>emit("update:modelValue",value)
})

const selected = () => {
  emit("selectd", activeCoverList.value);
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
