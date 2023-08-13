// composables/useFormatNumber.js
export default function useFormatNumber() {
  /**
   * 
   * @param {number} value 
   * @returns {string}
   * @description Format number to K or M
   */
  const formatNumber = (value) => {
    if (value >= 1e6) {
      return Math.floor(value / 1e6) + 'M';
    }
    if (value >= 1e3) {
      return Math.floor(value / 1e3) + 'K';
    }
    return value.toString();
  }

  return { formatNumber };
}
