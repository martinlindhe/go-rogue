// XXX from https://github.com/yuche/vue-strap/blob/master/src/utils/EventListener.js

export default {
    /**
     * Registers a DOM event listener to trigger during the bubble phase.
     *
     * @param {DOMEventTarget} target DOM element to register listener on.
     * @param {string} eventType Event type, e.g. 'click' or 'mouseover'.
     * @param {function} callback Callback function.
     * @return {object} Object with a `remove` method.
     */
    listen(target, eventType, callback) {
        if (target.addEventListener) {
            target.addEventListener(eventType, callback, false)
            return {
                remove() {
                    target.removeEventListener(eventType, callback, false)
                }
            }
        } else if (target.attachEvent) {
            target.attachEvent('on' + eventType, callback)
            return {
                remove() {
                    target.detachEvent('on' + eventType, callback)
                }
            }
        }
    }
};
