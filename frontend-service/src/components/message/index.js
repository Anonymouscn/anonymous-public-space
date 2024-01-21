import {createVNode, render} from "vue";
import Message from "./Message.vue";

// 向用户传递消息
const message = ({text, type, duration = 3000, customStyle = null}) => {
    if(document.getElementsByClassName('message-container').length > 0) {
        return
    }
    let point = document.getElementById('message-endpoint')
    if(!point) {
        point = document.body.createElement('div')
    }
    let timer = null
    const div = document.createElement('div')
    div.setAttribute('class', 'message-container')
    const vnode = createVNode(Message, {text, type, customStyle})
    render(vnode, div)
    const a = point.appendChild(div)
    timer = setTimeout(() => {
        setTimeout(() => {
            point.removeChild(a)
        }, duration * 1.5)
    })
}

export default message