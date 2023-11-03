// /composables/useSleep.js

export default function useSleep() {
    /**
     * 睡眠函数
     * @param {number} ms 毫秒
     */
    function sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
      }
    return { sleep };
}