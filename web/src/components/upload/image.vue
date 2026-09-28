
<template>
  <div>
    <el-upload
      ref="uploadRef"
      :show-file-list="false"
      :auto-upload="false"
      :on-change="handleChange"
      :multiple="multiple"
      action=""
    >
      <el-button type="primary">压缩上传</el-button>
    </el-upload>
  </div>
</template>

<script setup>
import ImageCompress from '@/utils/image'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import service from '@/utils/request'

defineOptions({
  name: 'UploadImage',
})

const emit = defineEmits(['on-success'])
const props = defineProps({
  imageUrl: {
    type: String,
    default: ''
  },
  fileSize: {
    type: Number,
    default: 2048 // 2M 超出后执行压缩
  },
  maxWH: {
    type: Number,
    default: 1920 // 图片长宽上限
  },
  multiple: {
    type: Boolean,
    default: false
  }
})

const uploadRef = ref(null)
const fileList = ref([])
let timer = null

const handleChange = (file) => {
  fileList.value.push(file)
  if (timer) clearTimeout(timer)
  timer = setTimeout(() => {
    handleUploadQueue()
  }, 300)
}

const handleUploadQueue = async() => {
  // 按照文件名从小到大排序
  const sortedFiles = [...fileList.value].sort((a, b) => {
    return a.name.localeCompare(b.name, undefined, { numeric: true, sensitivity: 'base' })
  })

  // 清理原始列表，避免重复执行
  fileList.value = []
  uploadRef.value.clearFiles()

  for (const fileItem of sortedFiles) {
    await uploadFile(fileItem.raw)
  }
}

const uploadFile = async(file) => {
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  const isWebp = file.type === 'image/webp'
  if (!isJPG && !isPng && !isWebp) {
    ElMessage.error('上传图片只能是 jpg, png 或 webp 格式!')
    return
  }

  let uploadFile = file
  const isRightSize = file.size / 1024 < props.fileSize
  if (!isRightSize) {
    // 压缩
    const compress = new ImageCompress(file, props.fileSize, props.maxWH)
    uploadFile = await compress.compress()
  }

  const formData = new FormData()
  formData.append('file', uploadFile)

  try {
    const res = await service({
      url: '/fileUploadAndDownload/upload',
      method: 'post',
      data: formData,
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    if (res.code === 0 && res.data.file) {
      emit('on-success', res.data.file.url)
    }
  } catch (error) {
    console.error('上传失败', error)
  }
}

</script>

<style lang="scss" scoped>
.image-uploader {
  border: 1px dashed #d9d9d9;
  width: 180px;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.image-uploader {
  border-color: #409eff;
}
.image-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.image {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
