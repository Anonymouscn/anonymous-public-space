<template>
  <!-- 页面容器 -->
  <div class="about-container lazy-background " ref="userView">
    <!-- 二维码弹窗 -->
    <pop-up
        :is-show="qrCodePopUpSettings.isShow"
        @close="closeQrCodePopUp"
        :show-header="false"
        :show-buttons="false"
        :width="qrCodePopUpSettings.width"
        :height="qrCodePopUpSettings.height"
    >
      <!-- 社交二维码内容插槽 -->
      <template v-slot:body>
        <div class="qr-code-box center--flex">
          <img :src="qrCodePopUpSettings.image" alt="#" class="qr-code" loading="lazy">
        </div>
      </template>
    </pop-up>
    <!-- 通知消息气泡 -->
    <div id="message-endpoint" ref="messageEndpoint" class="flex--center"></div>
    <div class="front-background-container lazy-background">
      <!-- seo 网站标题 -->
      <h1 class="seo">{{ seoSettings.site.title }}</h1>
      <!-- 个人介绍信 -->
      <div class="letter-box center--flex">
        <!-- 文本区 -->
        <div class="text-area center--flex" ref="mainWindow">
          <!-- 内容填充 -->
          <div class="text-area-content" >
            <p>Hi, guys 👋🏻.</p>
            <p>I'm anonymous.</p>
            <p>Good to see U in <span class="text-highlight">2024</span>.</p>
            <p>I work in <span class="text-highlight">Shenzhen</span>, <span class="text-highlight">China</span> as a Software Developer 🧑🏼‍💻.</p>
            <p>I have never had a photographic memory, so I create a <a class="text-link" @click="handleRequestToPlatform('#', 'SPACE')">SPACE</a> to record what I saw and heard in life 🌌.</p>
            <p>If you have any questions, you can leave me a <a class="text-link" @click="handleRequestToPlatform('#', 'MESSAGE')">MESSAGE</a> or <a class="text-link" @click="handleRequestToPlatform('#', 'CHAT')">CHAT</a> with my agent.</p>
            <p>Life is short, but digital are eternal !</p>
          </div>
          <!-- 联系方式 mini -->
          <Transition
              name="contact-m"
              @after-enter="afterContactMiniEnter"
              @after-leave="afterContactMiniLeave"
              @before-enter="beforeContactMiniEnter"
          >
            <div class="contact-mini" v-show="showMiniContact">
              <ContactList
                  :list="contactMini.list"
                  :item-size="contactMini.item.size"
                  :item-column-gap="contactMini.item.columnGap"
                  :item-row-gap="contactMini.item.rowGap"
                  ref="contactMiniList"
                  @action="openQrCodePopUp"
              />
            </div>
          </Transition>
        </div>
        <!-- 页脚信息 -->
        <footer class="footer-info center--flex" ref="footer">
          <!-- 备案信息 -->
          <div class="icp-licence">
            <span>{{ ICPLicence.province }}</span>
            <span>ICP</span>
            <span>备</span>
            <span>{{ ICPLicence.id }}</span>
            <span v-if="ICPLicence.no">-</span>
            <span v-if="ICPLicence.no">{{ ICPLicence.no }}</span>
            <span>号</span>
          </div>
          <!-- 版权信息 -->
          <div class="copyright-licence" v-if="copyrightLicence.author">
            <span>&copy;</span>
            <span>{{ copyrightLicence.from }}</span>
            <span v-if="copyrightLicence.from && copyrightLicence.to">-</span>
            <span>{{ copyrightLicence.to }}</span>
            <span>{{ copyrightLicence.author }}</span>
          </div>
        </footer>
      </div>
      <!-- 侧边栏 -->
      <Transition name="sidebar" @after-enter="afterSideBoxEnter" @after-leave="afterSideBoxLeave">
        <div class="side-box" ref="sideBox" v-show="showSideBox">
          <!-- 技术栈列表 -->
          <div class="intro-line" :ref="setSideBarTitle">Work with</div>
          <TechStackList
              :list="techStack.list"
              :item-size="techStack.itemSize"
              :item-column-gap="techStack.itemColumnGap"
              :item-row-gap="techStack.itemRowGap"
              ref="techStackList"
              class="side-box-item"
          />
          <!-- 社交联系方式列表 -->
          <div class="intro-line" :ref="setSideBarTitle">Contact me</div>
          <ContactList
              :list="contactMe.list"
              :item-size="contactMe.itemSize"
              :item-column-gap="contactMe.itemColumnGap"
              :item-row-gap="contactMe.itemRowGap"
              ref="contactList"
              @action="openQrCodePopUp"
              class="side-box-item"
          />
          <!-- 项目方式列表 -->
          <div class="intro-line" :ref="setSideBarTitle">Code together</div>
          <ContactList
              :list="codePlace.list"
              :item-size="codePlace.itemSize"
              :item-column-gap="codePlace.itemColumnGap"
              :item-row-gap="codePlace.itemRowGap"
              ref="codePlaceList"
              @action="openQrCodePopUp"
              class="side-box-item"
          />
          <!-- 站点入口列表 -->
          <div class="intro-line" :ref="setSideBarTitle">Go to</div>
          <div class="site-entrance">
            <a
                class="site-entrance-item side-box-item"
                v-for="item in siteEntrance.list"
                @click="handleRequestToPlatform(item.url, item.name)"
                :ref="setSiteEntrance"
            >
              {{ item.name }}
            </a>
          </div>
        </div>
      </Transition>
    </div>
  </div>
</template>

<script setup>
import {getCurrentInstance, onMounted, onUnmounted, reactive, ref, watch} from "vue";
import TechStackList from "../../components/TechStackList.vue";
import ContactList from "../../components/ContactList.vue";
import PopUp from "../../components/PopUp.vue";
import Message from "../../components/message/Message.vue";

// ========== 引用数据定义 ========== //

const mainWindow = ref(null) // mainWindow 引用
const sideBox = ref(null) // sideBox 引用
const techStackList = ref(null) // techStackList 引用
const contactList = ref(null) // contactList 引用
const codePlaceList = ref(null) // codeAtList 引用
const contactMiniList = ref(null) // contactMiniList 引用
const userView = ref(null) // userView 引用
const showSideBox = ref(false) // showSideBox 引用
const showMiniContact = ref(false) // showMiniContact 引用
const footer = ref(null) // footer 引用
const messageEndpoint = ref(null) // 消息切点
const sideBarTitle = ref([]) // 侧边栏标题 dom 数组
const letterContent = ref(null)
const setSideBarTitle = (el) => {
  sideBarTitle.value.push(el)
}
const siteEntranceList = ref([]) // 站点入口项 dom 数组
const setSiteEntrance = (el) => {
  siteEntranceList.value.push(el)
}
const { proxy } = getCurrentInstance() // vue 实例代理

// 自定义消息样式
const customMessage = reactive({
  style: {
    color: '#fff',
    backgroundColor: 'rgb(0, 0, 0, 0.3)',
    borderColor: 'rgb(60, 60, 60)',
    maxWidth: '86%',
    fontSize: '1.2rem',
    fontWeight: 500
  }
})

// ========== 后端数据定义 ========== //
const BaseUrl = import.meta.env.VITE_APP_BASE_URL || ''
// 背景加载路径
// const backgroundSource = ref('url(\''+ BaseUrl +'/background.output.webp\')')
const backgroundSource = ref('url(\'' +'/v1/background.output.webp\')')

// 备案信息
const ICPLicence = reactive({
  province: '粤',
  id: '12345678',
  no: '1'
})

// 版权信息
const copyrightLicence = reactive({
  from: '2021',
  to: '2024',
  author: 'anonymous'
})

// 技术栈列表
const techStack = reactive({
  list: [
    {icon: 'java', name: 'java'},
    {icon: 'spring', name: 'spring'},
    {icon: 'golang', name: 'golang'},
    {icon: 'html', name: 'html'},
    {icon: 'css3', name: 'css3'},
    {icon: 'javascript', name: 'javascript'},
    {icon: 'vue', name: 'vue'},
    {icon: 'linux', name: 'linux'},
    {icon: 'mysql', name: 'mysql'},
    {icon: 'redis', name: 'redis'},
    {icon: 'mongo', name: 'mongo'},
    {icon: 'docker', name: 'docker'},
    {icon: 'python', name: 'python'},
    {icon: 'pytorch', name: 'pytorch'}
  ],
  itemSize: '2.4rem',
  itemRowGap: '0.2rem',
  itemColumnGap: '0.2rem'
})

// 社交联系方式列表
const contactMe = reactive({
  list: [
    {icon: 'wechat', name: 'wechat', action: 'https://web-cache-1307314657.cos.ap-beijing.myqcloud.com/social-qrcode%2Fwechatme.webp'},
    {icon: 'qq', name: 'qq', action: 'https://web-cache-1307314657.cos.ap-beijing.myqcloud.com/social-qrcode%2Fqqme.webp'},
    {icon: 'telegram', name: 'telegram', url: 'https://t.me/anonymous0x10'}
  ],
  itemSize: '2.2rem',
  itemColumnGap: '0.4rem',
  itemRowGap: '0.3rem'
})

// 社交联系方式 mini 列表
const contactMini = reactive({
  list: [
    {icon: 'github', name: 'github', url: 'https://github.com/Anonymouscn/'},
    {icon: 'gitee', name: 'gitee', url: 'https://gitee.com/anonymous000/'},
    {icon: 'wechat', name: 'wechat', action: 'https://web-cache-1307314657.cos.ap-beijing.myqcloud.com/social-qrcode%2Fwechatme.webp'},
    {icon: 'qq', name: 'qq', action: 'https://web-cache-1307314657.cos.ap-beijing.myqcloud.com/social-qrcode%2Fqqme.webp'},
    {icon: 'telegram', name: 'telegram', url: 'https://t.me/anonymous0x10'}
  ],
  item: {
    size: '1.6rem',
    columnGap: '0.4rem',
    rowGap: '0.4rem'
  }
})

// 码农联系方式列表
const codePlace = reactive({
  list: [
    {icon: 'github', name: 'github', url: 'https://github.com/Anonymouscn/'},
    {icon: 'gitee', name: 'gitee', url: 'https://gitee.com/anonymous000/'}
  ],
  itemSize: '2.2rem',
  itemColumnGap: '0.4rem',
  itemRowGap: '0.3rem'
})

// 站点入口列表
const siteEntrance = reactive({
  list: [
    {name: 'Space', url: '#'},
    {name: 'Message', url: '#'},
    {name: 'Chat', url: '#'},
    {name: 'Manage', url: '$'}
  ]
})

// 二维码弹窗设置
const qrCodePopUpSettings = reactive({
  image: '#',
  width: '12rem',
  height: '12rem',
  isShow: false
})

// seo 信息设置
const seoSettings = reactive({
  site: {
    title: 'Anonymous Space',
    description: 'A space for summarizing my personal experiences, developing experiences, and making friends online, hoping to meet interesting souls!'
  }
})

// 打开二维码弹窗
const openQrCodePopUp = (src) => {
  qrCodePopUpSettings.image = src
  qrCodePopUpSettings.isShow = true
}

// 关闭二维码弹窗
const closeQrCodePopUp = () => {
  qrCodePopUpSettings.isShow = false
}

// userView 信息
const userViewInfo = reactive({
  width: 0,
  height: 0
})

// mainWindow 信息
const mainWindowInfo = reactive({
  width: 0,
  height: 0
})

// sideBox 信息
const sideBoxInfo = reactive({
  width: 0,
  height: 0
})

/*
 * 设置外观模式: 白天/黑夜
 * 1. 根据用户系统获取系统默认模式
 * 2. 系统默认模式获取失败: 尝试根据用户地理位置及其系统时间计算
 * 3. 获取用户地理位置失败: 尝试根据时间粗略推断 ([6:00, 18:00) = 白天，[18:00, 次日6:00) = 黑夜 )
 */
const setMode = () => {

}

// 处理个人介绍信内容
const produceLetter = () => {

}

// 处理前往平台请求
const handleRequestToPlatform = (url, name) => {
  if(!url || url === '') return
  // url='#' 站点正在维护升级
  if(url === '#') {
    const notice = 'platform is being upgraded or maintained, it will be coming soon'
    let text = name && name !== '' ? '[' + name + ']' + notice : notice;
    proxy.$message({text: text, type: 'custom', customStyle: customMessage.style})
    return;
  }
  // url='$' 站点无权限访问
  if(url === '$') {
    const notice = ' has no permission to access'
    let text = name && name !== '' ? 'platform [' + name + ']' + notice : 'platform ' + notice;
    proxy.$message({text: text})
    return;
  }
  // 站点跳转
  window.open(url)
}

// 背景延迟加载
const backgroundLazyLoad = () => {
  let lazyBackgrounds = [].slice.call(document.querySelectorAll(".lazy-background"))
  if ("IntersectionObserver" in window) {
    let lazyBackgroundObserver = new IntersectionObserver(function(entries, observer) {
      entries.forEach(function(entry) {
        if (entry.isIntersecting) {
          entry.target.classList.add("visible");
          lazyBackgroundObserver.unobserve(entry.target);
        }
      });
    });
    lazyBackgrounds.forEach(function(lazyBackground) {
      lazyBackgroundObserver.observe(lazyBackground);
    });
  }
}

// 获取 userView 信息
const getUserViewInfo = () => {
  userViewInfo.width = userView.value.offsetWidth
  userViewInfo.height = userView.value.offsetHeight
}

// 获取 mainWindow 信息
const getMainWindowInfo = () => {
  mainWindowInfo.width = mainWindow.value.offsetWidth
  mainWindowInfo.height = mainWindow.value.offsetHeight
}

// 获取 sideBox 信息
const getSideBoxInfo = () => {
  sideBoxInfo.width = sideBox.value.offsetWidth
  sideBoxInfo.height = sideBox.value.offsetHeight
}

// 通知 sideBox 大小自适应
const noticeSideBoxInfo = () => {
  resetTechTaskList() // 重置 techStackList 大小
  resetContactMeList() // 重置 contactMeList 大小
  resetCodePlaceList() // 重置 codePlaceList 大小
}
watch(sideBoxInfo, noticeSideBoxInfo)

// 计算 sideBox 列表最大行数
const calculateSideBoxListMaxRows = () => {
  let height = sideBoxInfo.height
  return height < 850 ? (height < 625 ? 1 : 2) : 3
}

// sideBox 右侧留白填充大小
const sideBoxPaddingRight = 5

// 重置 techStackList 大小
const resetTechTaskList = () => {
  let maxRow = calculateSideBoxListMaxRows()
  let width = sideBoxInfo.width
  let itemSize = techStack.itemSize
  if(!itemSize) itemSize = '0rem'
  let itemGap = techStack.itemColumnGap
  if(!itemGap) itemGap = '0rem'
  itemSize = parseInt(parseFloat(itemSize.substr(0, itemSize.length - 3) + 0.4) * 16)
  itemGap = parseInt(parseFloat(itemGap.substr(0, itemGap.length - 3)) * 16)
  let padding = 32
  let len = parseInt((width - padding) / (itemSize + itemGap + sideBoxPaddingRight))
  techStackList.value.resize(len, maxRow)
}

// 重置 contactList 大小
const resetContactMeList = () => {
  let maxRow = calculateSideBoxListMaxRows()
  let width = sideBoxInfo.width
  let itemSize = contactMe.itemSize
  if(!itemSize) itemSize = '0rem'
  let itemGap = contactMe.itemColumnGap
  if(!itemGap) itemGap = '0rem'
  itemSize = parseInt(parseFloat(itemSize.substr(0, itemSize.length - 3)) * 16)
  itemGap = parseInt(parseFloat(itemGap.substr(0, itemGap.length - 3) + 0.4) * 16)
  let padding = 32
  let len = parseInt((width - padding) / (itemSize + itemGap + sideBoxPaddingRight))
  contactList.value.resize(len, maxRow)
}

// 重置 codePlaceList 大小
const resetCodePlaceList = () => {
  let maxRow = calculateSideBoxListMaxRows()
  let width = sideBoxInfo.width
  let itemSize = codePlace.itemSize
  if(!itemSize) itemSize = '0rem'
  let itemGap = codePlace.itemColumnGap
  if(!itemGap) itemGap = '0rem'
  itemSize = parseInt(parseFloat(itemSize.substr(0, itemSize.length - 3)) * 16)
  itemGap = parseInt(parseFloat(itemGap.substr(0, itemGap.length - 3) + 0.4) * 16)
  let padding = 32
  let len = parseInt((width - padding) / (itemSize + itemGap + sideBoxPaddingRight))
  codePlaceList.value.resize(len, maxRow)
}

// 通知 mainWindow 大小自适应
const noticeMainWindowInfo = () => {
  resizeContactMiniList() // 重置 contactMiniList 大小
}
watch(mainWindowInfo, noticeMainWindowInfo)

// 重置 contactMiniList 大小
const resizeContactMiniList = () => {
  let width = mainWindowInfo.width
  let itemSize = contactMini.item.size
  if(!itemSize) itemSize = '0rem'
  let itemGap = contactMini.item.gap
  if(!itemGap) itemGap = '0rem'
  itemSize = parseInt(parseFloat(itemSize.substr(0, itemSize.length - 3)) * 16)
  itemGap = parseInt(parseFloat(itemGap.substr(0, itemGap.length - 3)) * 16)
  let max = parseInt(width / (itemSize + itemGap))
  let len = contactMini.list.length
  len = len > max ? max : len
  contactMiniList.value.resize(len, 1)
}

// 通知 userView 大小自适应
const noticeUserViewInfo = () => {
  // 获取 userView 大小
  showSideBox.value = userViewInfo.width > 960
  showMiniContact.value = userViewInfo.width <= 960
  let vw = userViewInfo.width
  letterFontAutoFit(vw) // 个人介绍信字体自适应
  sideBarTitleFontAutoFit(vw) // 侧边栏标题字体自适应
  siteEntranceFontAutoFit(vw) // 站点入口字体自适应
  footerInfoFontAutoFit(vw) // 页脚字体自适应
  messageFontAutoFit(vw)
}
watch(userViewInfo, noticeUserViewInfo)

// ========== 字体设置 ========== //

// 个人介绍信字体自适应
const letterFontAutoFit = (width) => {
  let cls = mainWindow.value.classList
  if(width < 360) { // mini phone
    resetClassList(cls, 'phone-mini-text-area')
  } else if(width < 430) { // normal phone
    resetClassList(cls, 'phone-normal-text-area')
  } else if(width < 512) { // iphone 14 pro max
    resetClassList(cls, 'phone-14pm-text-area')
  } else if(width < 768) { // mini pad
    resetClassList(cls, 'pad-mini-text-area')
  } else if(width < 1280) { // normal pad
    resetClassList(cls, 'pad-normal-text-area')
  } else if(width < 1440) { // pc
    resetClassList(cls, 'pc-text-area')
  } else if(width < 1720) { // pc medium
    resetClassList(cls, 'pc-medium-text-area')
  } else { // pc large
    resetClassList(cls, 'pc-large-text-area')
  }
}

// 侧边栏标题字体自适应
const sideBarTitleFontAutoFit = (width) => {
  let titles = sideBarTitle.value
  for(let i = 0; i < titles.length; ++i) {
    let cls = titles[i].classList
    if(width < 1280) { // phone or pad
      resetClassList(cls, 'pad-normal-title')
    } else { // pc
      resetClassList(cls, 'pc-normal-title')
    }
  }
}

// 站点入口字体自适应
const siteEntranceFontAutoFit = (width) => {
  let sites = siteEntranceList.value
  for(let i = 0; i < sites.length; ++i) {
    let cls = sites[i].classList
    if(width < 1280) { // phone or pad
      resetClassList(cls, 'pad-normal-site-entrance-font')
    } else if(width < 1440) { // ipad pro
      resetClassList(cls, 'pad-ipad-pro-site-entrance-font')
    } else { // pc
      resetClassList(cls, 'pc-pro-site-entrance-font')
    }
  }
}

// 页脚字体自适应
const footerInfoFontAutoFit = (width) => {
  let cls = footer.value.classList
  if(width < 768) { // phone
    resetClassList(cls, 'footer-info-phone')
  } else if(width < 1280) { // pad
    resetClassList(cls, 'footer-info-pad')
  } else { // pc
    resetClassList(cls, 'footer-info-pc')
  }
}

// 消息字体自适应
const messageFontAutoFit = (width) => {
  let s = customMessage.style
  if(width < 360) { // mini phone
    s.fontSize = '0.8rem'
    s.lineHeight = '1.2rem'
    s.padding = '0.2rem 1rem'
  } else if(width < 430) { // normal phone
    s.fontSize = '0.9rem'
    s.lineHeight = '1.3rem'
    s.padding = '0.2rem 1rem'
  } else if(width < 512) { // iphone 14 pro max
    s.fontSize = '1rem'
    s.lineHeight = '1.4rem'
    s.padding = '0.2rem 1rem'
  } else if(width < 768) { // mini pad
    s.fontSize = '1.1rem'
    s.lineHeight = '1.5rem'
    s.padding = '0.2rem 1rem'
  } else if(width < 1280) { // normal pad
    s.fontSize = '1.2rem'
    s.lineHeight = '1.6rem'
    s.padding = '0.2rem 1rem'
  } else  { // pc
    s.fontSize = '1.3rem'
    s.lineHeight = '1.6rem'
    s.padding = '0.4rem 1rem'
  }
}

// 重置样式表
const resetClassList = (cls, add) => {
  for (let i = 0; i < cls.length; ) {
    if (cls[i].match('phone') || cls[i].match('pad') || cls[i].match('pc')) {
      cls.remove(cls[i])
    } else {
      ++i
    }
  }
  cls.add(add)
}

// ========== 动画设置 ========== //

// sideBox 完成渲染后获取宽高
const afterSideBoxEnter = () => {
  getSideBoxInfo()
}

// sideBox 移除渲染后获取宽高
const afterSideBoxLeave = () => {
  getMainWindowInfo()
}

// contactMini 完成渲染前获取宽高
const beforeContactMiniEnter = () => {
  getMainWindowInfo()
}

// contactMini 完成渲染后获取宽高
const afterContactMiniEnter = () => {
  getMainWindowInfo()
}

// contactMini 移除渲染后获取宽高
const afterContactMiniLeave = () => {
  getMainWindowInfo()
}

// ========== 函数钩子设置 ========== //

// 挂载钩子函数
onMounted(() => {
  getUserViewInfo()
  window.addEventListener("resize", getUserViewInfo)
  window.addEventListener("resize", getMainWindowInfo)
  window.addEventListener("resize", getSideBoxInfo)
  window.onload = backgroundLazyLoad()
})

// 卸载钩子函数
onUnmounted(() => {
  window.removeEventListener("resize", getUserViewInfo)
  window.removeEventListener("resize", getMainWindowInfo)
  window.removeEventListener("resize", getSideBoxInfo)
})
</script>

<style scoped>
.message-endpoint {
  position: absolute;
  z-index: 3;
  display: flex;
  flex-direction: column;
}

/* 页面容器 */
.about-container {
  position: absolute;
  left: 0;
  top: 0;
  height: 100%;
  font-family: SongTi, sans-serif;
  background-image: v-bind(backgroundSource);
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
  background-position: bottom;
  overflow: hidden;
}

/* 前置背景容器 */
.front-background-container {
  position: relative;
  margin: auto;
  width: 100vw;
  height: 100%;
  color: #fff;
  font-size: 30px;
  background-color: rgba(255, 255, 255, 0.2);
  overflow: hidden;
  z-index: 0;
  display: flex;
}

.front-background-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100%;
  filter: blur(18px);
  background-image: v-bind(backgroundSource);
  background-repeat: no-repeat;
  background-attachment: fixed;
  background-size: cover;
  background-position: bottom;
}

.front-background-container::after {
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  content: '';
}

.dark .front-background-container::after {
  background-color: rgb(22, 22, 22, 0.8);
}

/* 介绍信 */
.letter-box {
  flex: 1;
  overflow: auto;
  z-index: 1;
}

/* 文本区 */
.text-area {
  width: 80%;
  flex-direction: column;
  align-items: flex-start;
  word-spacing: 2px;
  word-break: break-word;
  max-width: 55rem;
  color: rgb(245, 245, 245, 0.9);
}

/* 文本区段落 */
.text-area p {
  font-family: SongTi, serif;
  font-weight: 600;
}

/* 文本区段落 - phone mini */
.phone-mini-text-area p {
  font-size: 0.8rem;
  line-height: 1.8rem;
}

/* 文本区段落 - phone normal */
.phone-normal-text-area p {
  font-size: 1rem;
  line-height: 2.2rem;
}

/* 文本区段落 - iphone 14 pro max */
.phone-14pm-text-area p {
  font-size: 1.2rem;
  line-height: 2.6rem;
}

/* 文本区段落 - pad mini */
.pad-mini-text-area p {
  font-size: 1.2rem;
  line-height: 2.6rem;
}

/* 文本区段落 - pad normal */
.pad-normal-text-area p {
  font-size: 1.3rem;
  line-height: 2.7rem;
}

/* 文本区段落 - pc */
.pc-text-area p {
  font-size: 1.4rem;
  line-height: 3rem;
}

/* 文本区段落 - pc medium */
.pc-medium-text-area p {
  font-size: 1.5rem;
  line-height: 3.2rem;
}

/* 文本区段落 - pc large */
.pc-large-text-area p {
  font-size: 1.7rem;
  line-height: 3.6rem;
}

/* 文本高亮 */
.text-highlight {
  font-weight: 800;
}

/* 文本高亮 - phone mini */
.phone-mini-text-area .text-highlight {
  font-size: 0.9rem;
}

/* 文本高亮 - phone normal */
.phone-normal-text-area .text-highlight {
  font-size: 1.1rem;
}

/* 文本高亮 - iphone 14 pro max */
.phone-14pm-text-area .text-highlight {
  font-size: 1.3rem;
}

/* 文本高亮 - pad mini */
.pad-mini-text-area .text-highlight {
  font-size: 1.4rem;
}

/* 文本高亮 - pc */
.pc-text-area .text-highlight {
  font-size: 1.5rem;
}

/* 文本高亮 - pc medium */
.pc-medium-text-area .text-highlight {
  font-size: 1.6rem;
}

/* 文本高亮 - pc large */
.pc-large-text-area .text-highlight {
  font-size: 1.8rem;
}

/* 文本链接 */
.text-link {
  color: rgb(250, 250, 250, 0.8);
  text-decoration: none;
  position: relative;
  font-weight: 700;
  transition: color ease .5s;
  cursor: pointer;
}

.text-link::after {
  content: '';
  position: absolute;
  bottom: -0.2rem;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: rgb(250, 250, 250, 0.8);
  transition: background-color ease .5s;
}

@media (any-hover: hover) {
  .text-link:hover {
    color: #fff;
  }

  .text-link:hover.text-link::after {
    background-color: #fff;
  }
}

/* 文本链接 - phone mini */
.phone-mini-text-area .text-link {
  font-size: 0.8rem;
}

/* 文本链接 - phone normal */
.phone-normal-text-area .text-link {
  font-size: 1rem;
}

/* 文本链接 - iphone 14 pro max */
.phone-14pm-text-area .text-link {
  font-size: 1.2rem;
}

/* 文本链接 - pad mini */
.pad-mini-text-area .text-link {
  font-size: 1.2rem;
}

/* 文本链接 - pc */
.pc-text-area .text-link {
  font-size: 1.3rem;
}

/* 文本链接 - pc medium */
.pc-medium-text-area .text-link {
  font-size: 1.4rem;
}

/* 文本链接 - pc large */
.pc-large-text-area .text-link {
  font-size: 1.5rem;
}

/* 文本区内容 */
.text-area-content {
  transition: all 0.5s ease-in-out;
}

/* mini 联系方式列表 */
.contact-mini {
  align-self: flex-end;
  min-height: 2rem;
  margin-top: 0.5rem;
}

/* 页脚信息 */
.footer-info {
  position: absolute;
  bottom: 0;
  font-size: 0;
  color: #495057;
  padding-bottom: 0.5rem;
  gap: 0;
  text-align: center;
  transition: all ease 0.3s;
}

/* 页脚信息 - phone */
.footer-info-phone {
  font-size: 0.6rem;
  gap: 0.3rem;
}

/* 页脚信息 - pad */
.footer-info-pad {
  font-size: 0.7rem;
  gap: 0.4rem;
}

/* 页脚信息 - pc */
.footer-info-pc {
  font-size: 0.8rem;
  gap: 0.5rem;
}

/* ICP 信息 */
.icp-licence {
  display: inline-flex;
  align-items: center;
  justify-items: center;
  gap: 2px;
}

/* 版权信息 */
.copyright-licence {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 2px;
}

/* 侧边栏 */
.side-box {
  margin-left: auto;
  width: 26%;
  min-width: 18rem;
  max-width: 28rem;
  height: 100vh;
  font-size: 30px;
  background-color: rgba(174, 169, 164, 0.5);
  overflow: hidden;
  z-index: 2;
  box-shadow: 0 3px 6px #959595;
  display: flex;
  flex-direction: column;
  padding: 1rem 2rem;
  gap: 0.5rem;
}

.dark .side-box {
  background-color: rgba(174, 169, 164, 0.2);
}

@media screen and (orientation: portrait) {
  .side-box {
    display: none;
  }
  .contact-mini {
    display: flex !important;
  }
}

.intro-line {
  font-weight: 700;
  color: rgb(250, 250, 250);
}

.pad-normal-title {
  font-size: 1.6rem;
  line-height: 2.8rem;
}

.pc-normal-title {
  font-size: 1.7rem;
  line-height: 2.8rem;
}

.side-box-item {
  margin-bottom: 0.5rem;
}

.side-box-item:last-child {
  margin-bottom: 0;
}

/* 站点入口列表 */
.site-entrance {
  word-wrap: break-word;
  line-height: 1.8rem;
}

/* 站点入口列表项 */
.site-entrance-item {
  font-size: 1.6rem;
  line-height: 1.8rem;
  text-decoration: none;
  color: rgb(250, 250, 250, 0.8);
  display: inline-block;
  margin-right: 1rem;
  margin-bottom: 0.5rem;
  position: relative;
  transition: color ease .5s;
  cursor: pointer;
}

@media (any-hover: hover) {
  .site-entrance-item:hover {
    color: #fff;
  }

  .site-entrance-item:hover.side-box-item::after {
    background-color: #fff;
  }
}

.site-entrance-item:last-child {
  margin-right: 0;
}

.side-box-item::after {
  content: '';
  position: absolute;
  bottom: 0;
  transform: translateY(2px);
  height: 2px;
  width: 100%;
  background-color: rgb(250, 250, 250, 0.8);
  left: 0;
  transition: background-color ease .5s;
}

.pad-normal-site-entrance-font {
  font-size: 1.2rem;
  line-height: 1.5rem;
}

.pad-ipad-pro-site-entrance-font {
  font-size: 1.3rem;
  line-height: 1.6rem;
}

.pc-pro-site-entrance-font {
  font-size: 1.4rem;
  line-height: 1.7rem;
}

/* 二维码弹窗 */
.qr-code-box {
  width: 100%;
  height: 100%;
}

/* 二维码图片 */
.qr-code {
  width: 85%;
  height: 85%;
}

/* 侧边栏动画 */
.sidebar-enter-active {
  transition: all 0.5s ease-out;
}

.sidebar-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}

.sidebar-enter-from,
.sidebar-leave-to {
  transform: translateX(20px);
  opacity: 0;
}

/* mini 联系列表动画 */
.contact-m-enter-active {
  animation: bounce-in 0.5s;
}

.contact-m-leave-active {
  animation: bounce-in 0.5s reverse;
}

@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.25);
  }
  100% {
    transform: scale(1);
  }
}

/* 仅网站 seo 作用元素 */
.seo {
  display: none;
}

/* 消息弹窗切点 */
#message-endpoint {
  position: absolute;
  left: 0;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  top: 1.5rem;
  z-index: 1;
  flex-direction: column;
}
</style>