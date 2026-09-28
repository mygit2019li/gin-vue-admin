<template>
  <div>
    <div class="gva-table-box">
      <el-tabs v-model="activeTab">
        <!-- 分类管理 -->
        <el-tab-pane label="分类管理" name="category">
          <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openCategoryDialog">新增分类</el-button>
          </div>
          <el-table :data="categories" style="width: 100%" row-key="ID">
            <el-table-column align="left" label="排序" prop="sort" width="80" />
            <el-table-column align="left" label="分类名称" prop="name" min-width="120" />
            <el-table-column align="left" label="状态" width="100">
              <template #default="scope">
                <el-switch
                  v-model="scope.row.isHidden"
                  :active-value="false"
                  :inactive-value="true"
                  active-text="显示"
                  inactive-text="隐藏"
                  inline-prompt
                  @change="updateCategoryStatus(scope.row)"
                />
              </template>
            </el-table-column>
            <el-table-column align="left" label="操作" width="160">
              <template #default="scope">
                <el-button type="primary" link icon="edit" @click="editCategoryAction(scope.row)">编辑</el-button>
                <el-popover v-model="scope.row.visible" placement="top" width="160">
                  <p>确定要删除吗？</p>
                  <div style="text-align: right; margin-top: 8px;">
                    <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                    <el-button type="primary" @click="deleteCategoryAction(scope.row)">确定</el-button>
                  </div>
                  <template #reference>
                    <el-button type="primary" link icon="delete" @click="scope.row.visible = true">删除</el-button>
                  </template>
                </el-popover>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>

        <!-- 文案管理 -->
        <el-tab-pane label="文案管理" name="copywriting">
          <div class="gva-search-box">
            <el-form :inline="true" :model="searchInfo">
              <el-form-item label="分类">
                <el-select v-model="searchInfo.categoryId" placeholder="请选择分类" clearable @change="getCopywritingData">
                  <el-option v-for="item in categories" :key="item.ID" :label="item.name" :value="item.ID" />
                </el-select>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" icon="search" @click="getCopywritingData">查询</el-button>
              </el-form-item>
            </el-form>
          </div>
          <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openCopywritingDialog">新增文案</el-button>
          </div>
          <el-table :data="copywritingList" style="width: 100%" row-key="ID">
            <el-table-column align="left" label="排序" prop="sort" width="80" />
            <el-table-column align="left" label="文案内容" prop="text" min-width="200" show-overflow-tooltip />
            <el-table-column align="left" label="图片预览" width="320">
              <template #default="scope">
                <div class="table-images-grid">
                  <div v-for="(img, index) in (scope.row.images ? JSON.parse(scope.row.images) : []).slice(0, 9)" :key="index" class="table-img-item">
                    <CustomPic pic-type="file" :pic-src="img" preview fit="cover" />
                  </div>
                </div>
              </template>
            </el-table-column>
            <el-table-column align="left" label="状态" width="100">
              <template #default="scope">
                <el-switch
                  v-model="scope.row.isHidden"
                  :active-value="false"
                  :inactive-value="true"
                  active-text="显示"
                  inactive-text="隐藏"
                  inline-prompt
                  @change="updateCopywritingStatus(scope.row)"
                />
              </template>
            </el-table-column>
            <el-table-column align="left" label="操作" width="160">
              <template #default="scope">
                <el-button type="primary" link icon="edit" @click="editCopywritingAction(scope.row)">编辑</el-button>
                <el-popover v-model="scope.row.visible" placement="top" width="160">
                  <p>确定要删除吗？</p>
                  <div style="text-align: right; margin-top: 8px;">
                    <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                    <el-button type="primary" @click="deleteCopywritingAction(scope.row)">确定</el-button>
                  </div>
                  <template #reference>
                    <el-button type="primary" link icon="delete" @click="scope.row.visible = true">删除</el-button>
                  </template>
                </el-popover>
              </template>
            </el-table-column>
          </el-table>
          <div class="gva-pagination">
            <el-pagination
              :current-page="page"
              :page-size="pageSize"
              :page-sizes="[10, 30, 50, 100]"
              :total="total"
              layout="total, sizes, prev, pager, next, jumper"
              @current-change="handleCurrentChange"
              @size-change="handleSizeChange"
            />
          </div>
        </el-tab-pane>
      </el-tabs>
    </div>

    <!-- 分类对话框 -->
    <el-dialog v-model="categoryDialogVisible" :title="categoryType==='create'?'新增分类':'编辑分类'">
      <el-form :model="categoryForm" label-width="80px">
        <el-form-item label="名称">
          <el-input v-model="categoryForm.name" autocomplete="off" />
        </el-form-item>
        <el-form-item label="状态">
          <el-switch
            v-model="categoryForm.isHidden"
            :active-value="false"
            :inactive-value="true"
            active-text="显示"
            inactive-text="隐藏"
            inline-prompt
          />
        </el-form-item>
        <el-form-item label="排序">
          <el-input-number v-model="categoryForm.sort" :min="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="categoryDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="enterCategoryDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 文案对话框 -->
    <el-dialog v-model="cpDialogVisible" :title="cpType==='create'?'新增文案':'编辑文案'" width="60%">
      <el-form :model="cpForm" label-width="80px">
        <el-form-item label="分类">
          <el-select v-model="cpForm.categoryId" placeholder="请选择分类">
            <el-option v-for="item in categories" :key="item.ID" :label="item.name" :value="item.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="文案内容">
          <el-input v-model="cpForm.text" type="textarea" :rows="4" placeholder="请输入朋友圈文案" />
        </el-form-item>
        <el-form-item label="图片">
          <div class="image-list">
            <div
              v-for="(img, index) in cpFormImages"
              :key="index"
              class="image-item"
              draggable="true"
              @dragstart="handleDragStart(index)"
              @dragover.prevent="handleDragOver(index)"
              @drop="handleDrop(index)"
            >
              <CustomPic pic-type="file" :pic-src="img" />
              <el-icon class="delete-icon" @click="removeImage(index)"><CircleClose /></el-icon>
            </div>
            <upload-image :multiple="true" @on-success="addImage" />
          </div>
          <div v-if="cpFormImages.length > 0" class="image-tip-box">
            <span class="image-tip">共 {{ cpFormImages.length }} 张图片 (最多 9 张)</span>
            <el-button type="primary" link @click="previewGrid">预览九宫格</el-button>
          </div>
        </el-form-item>
        <el-form-item label="状态">
          <el-switch
            v-model="cpForm.isHidden"
            :active-value="false"
            :inactive-value="true"
            active-text="显示"
            inactive-text="隐藏"
            inline-prompt
          />
        </el-form-item>
        <el-form-item label="排序">
          <el-input-number v-model="cpForm.sort" :min="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="cpDialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="enterCopywritingDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 九宫格预览 -->
    <el-dialog v-model="gridVisible" title="九宫格预览" width="360px" custom-class="grid-preview-dialog">
      <div class="nine-grid">
        <div v-for="(img, index) in 9" :key="index" class="grid-item">
          <CustomPic v-if="cpFormImages[index]" pic-type="file" :pic-src="cpFormImages[index]" />
          <div v-else class="grid-placeholder"></div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createCategory,
  deleteCategory,
  updateCategory,
  getCategoryList,
  createCopywriting,
  deleteCopywriting,
  updateCopywriting,
  getCopywritingList
} from '@/api/miniProgram'
import UploadImage from '@/components/upload/image.vue'
import CustomPic from '@/components/customPic/index.vue'
import { CircleClose } from '@element-plus/icons-vue'
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'MomentsManage'
})

const activeTab = ref('category')
const categories = ref([])
const copywritingList = ref([])
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const searchInfo = ref({ categoryId: undefined })

// 获取分类
const getCategories = async() => {
  const res = await getCategoryList()
  if (res.code === 0) {
    categories.value = res.data.list
  }
}
getCategories()

// 获取文案
const getCopywritingData = async() => {
  const res = await getCopywritingList({
    page: page.value,
    pageSize: pageSize.value,
    categoryId: searchInfo.value.categoryId
  })
  if (res.code === 0) {
    copywritingList.value = res.data.list
    total.value = res.data.total
    page.value = res.data.page
    pageSize.value = res.data.pageSize
  }
}

watch(activeTab, (val) => {
  if (val === 'copywriting') {
    getCopywritingData()
  }
})

const handleSizeChange = (val) => {
  pageSize.value = val
  getCopywritingData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getCopywritingData()
}

// 分类对话框逻辑
const categoryDialogVisible = ref(false)
const categoryType = ref('create')
const categoryForm = ref({ name: '', sort: 0 })

const openCategoryDialog = () => {
  categoryType.value = 'create'
  categoryForm.value = { name: '', sort: 0 }
  categoryDialogVisible.value = true
}

const editCategoryAction = (row) => {
  categoryType.value = 'update'
  categoryForm.value = { ...row }
  categoryDialogVisible.value = true
}

const deleteCategoryAction = async(row) => {
  row.visible = false
  const res = await deleteCategory({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    getCategories()
  }
}

const enterCategoryDialog = async() => {
  let res
  if (categoryType.value === 'create') {
    res = await createCategory(categoryForm.value)
  } else {
    res = await updateCategory(categoryForm.value)
  }
  if (res.code === 0) {
    if (categoryType.value === 'create') {
      ElMessage.success('创建成功')
    } else {
      ElMessage.success('更新成功')
    }
    categoryDialogVisible.value = false
    getCategories()
  }
}

const updateCategoryStatus = async(row) => {
  const res = await updateCategory(row)
  if (res.code === 0) {
    ElMessage.success('状态更新成功')
  } else {
    row.isHidden = !row.isHidden
  }
}

// 文案对话框逻辑
const cpDialogVisible = ref(false)
const gridVisible = ref(false)
const cpType = ref('create')
const cpForm = ref({ categoryId: undefined, text: '', images: '[]', sort: 0 })
const cpFormImages = ref([])
const dragIndex = ref(null)

const handleDragStart = (index) => {
  dragIndex.value = index
}

const handleDragOver = (index) => {
  // 仅为了允许放置
}

const handleDrop = (index) => {
  if (dragIndex.value === null) return
  const item = cpFormImages.value.splice(dragIndex.value, 1)[0]
  cpFormImages.value.splice(index, 0, item)
  dragIndex.value = null
}

const previewGrid = () => {
  gridVisible.value = true
}

const openCopywritingDialog = () => {
  cpType.value = 'create'
  cpForm.value = { categoryId: categories.value[0]?.ID, text: '', images: '[]', sort: 0 }
  cpFormImages.value = []
  cpDialogVisible.value = true
}

const editCopywritingAction = (row) => {
  cpType.value = 'update'
  cpForm.value = { ...row }
  cpFormImages.value = row.images ? JSON.parse(row.images) : []
  cpDialogVisible.value = true
}

const addImage = (url) => {
  if (cpFormImages.value.length >= 9) {
    ElMessage.warning('最多只能上传 9 张图片')
    return
  }
  cpFormImages.value.push(url)
}

const removeImage = (index) => {
  cpFormImages.value.splice(index, 1)
}

const deleteCopywritingAction = async(row) => {
  row.visible = false
  const res = await deleteCopywriting({ ID: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    getCopywritingData()
  }
}

const enterCopywritingDialog = async() => {
  cpForm.value.images = JSON.stringify(cpFormImages.value)
  let res
  if (cpType.value === 'create') {
    res = await createCopywriting(cpForm.value)
  } else {
    res = await updateCopywriting(cpForm.value)
  }
  if (res.code === 0) {
    ElMessage.success('保存成功')
    cpDialogVisible.value = false
    getCopywritingData()
  }
}

const updateCopywritingStatus = async(row) => {
  const res = await updateCopywriting(row)
  if (res.code === 0) {
    ElMessage.success('状态更新成功')
  } else {
    // Revert if failed
    row.isHidden = !row.isHidden
  }
}
</script>

<style scoped lang="scss">
.image-list {
  display: grid;
  grid-template-columns: repeat(3, 80px);
  gap: 10px;
  .image-item {
    position: relative;
    width: 80px;
    height: 80px;
    cursor: grab;
    &:active {
      cursor: grabbing;
    }
    .delete-icon {
      position: absolute;
      top: -5px;
      right: -5px;
      color: #f56c6c;
      cursor: pointer;
      font-size: 18px;
      z-index: 10;
    }
  }
}
.image-tip-box {
  display: flex;
  align-items: center;
  gap: 15px;
  margin-top: 5px;
}
.image-tip {
  font-size: 12px;
  color: #909399;
}

.nine-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 8px;
  width: 300px;
  margin: 0 auto;
  .grid-item {
    width: 90px;
    height: 90px;
    background: #f5f7fa;
    border-radius: 4px;
    overflow: hidden;
    .grid-placeholder {
      width: 100%;
      height: 100%;
      background: #f0f2f5;
    }
  }
}

.table-images-grid {
  display: grid;
  grid-template-columns: repeat(3, 40px);
  gap: 0;
  width: fit-content;
  .table-img-item {
    width: 40px;
    height: 40px;
    border-radius: 0;
    overflow: hidden;
  }
}
</style>
