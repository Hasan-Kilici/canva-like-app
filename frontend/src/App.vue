<template>
  <div class="container">
      <div class="sidebar">
      <input class="search" v-model="searchTerm" placeholder="Enter search term">
      <button class="search-btn" @click="searchImages">Search Images</button>

      <div style="display:flex;align-items:stretch;flex-wrap:wrap;" v-if="images.length > 0">
          <div v-for="image in images" :key="image.id">
            <button @click="editor.chain().focus().setImage({ src: image.urls.small }).run()" style="padding:10px;background-color:transparent;"><img height="100px" :src="image.urls.small" alt="Unsplash Image"></button>
          </div>
      </div>
    </div>
    <div class="content">
    <button @click="editor.chain().focus().toggleHeading({ level: 1 }).run()">
          H1
        </button>
        <button @click="editor.chain().focus().toggleHeading({ level: 2 }).run()">
          H2
        </button>
        <button @click="editor.chain().focus().toggleHeading({ level: 3 }).run()">
          H3
        </button>
        <button @click="editor.chain().focus().toggleHeading({ level: 4 }).run()">
          H4
        </button>
        <button @click="editor.chain().focus().toggleHeading({ level: 5 }).run()">
          H5
        </button>
        <button @click="editor.chain().focus().toggleHeading({ level: 6 }).run()">
          H6
        </button>
        <button @click="editor.chain().focus().toggleBulletList().run()">
          Bullet List
        </button>
    <div v-if="editor">
      <bubble-menu
        class="bubble-menu"
        :tippy-options="{ duration: 100 }"
        :editor="editor"
      >
        <button @click="editor.chain().focus().toggleBold().run()" :class="{ 'is-active': editor.isActive('bold') }">
          Bold
        </button>
        <button @click="editor.chain().focus().toggleItalic().run()" :class="{ 'is-active': editor.isActive('italic') }">
          Italic
        </button>
        <button @click="editor.chain().focus().toggleStrike().run()" :class="{ 'is-active': editor.isActive('strike') }">
          Strike
        </button>
        <button @click="editor.chain().focus().toggleUnderline().run()" :class="{ 'is-active': editor.isActive('strike') }">
          Underline
        </button>
        <button @click="editor.chain().focus().setTextAlign('left').run()" :class="{ 'is-active': editor.isActive({ textAlign: 'left' }) }">
      left
    </button>
    <button @click="editor.chain().focus().setTextAlign('center').run()" :class="{ 'is-active': editor.isActive({ textAlign: 'center' }) }">
      center
    </button>
    <button @click="editor.chain().focus().setTextAlign('right').run()" :class="{ 'is-active': editor.isActive({ textAlign: 'right' }) }">
      right
    </button>
    <button @click="editor.chain().focus().setTextAlign('justify').run()" :class="{ 'is-active': editor.isActive({ textAlign: 'justify' }) }">
      justify
    </button>
        <div class="colors">
          <button class="color red" @click="editor.chain().focus().setColor('red').run()" :class="{ 'is-active': editor.isActive('textStyle', { color: 'red' })}"></button>
          <button class="color blue" @click="editor.chain().focus().setColor('blue').run()" :class="{ 'is-active': editor.isActive('textStyle', { color: 'blue' })}"></button>
          <button class="color green" @click="editor.chain().focus().setColor('green').run()" :class="{ 'is-active': editor.isActive('textStyle', { color: 'green' })}"></button>
          <button class="color yellow" @click="editor.chain().focus().setColor('yellow').run()" :class="{ 'is-active': editor.isActive('textStyle', { color: 'yellow' })}"></button>
          <button class="color lime" @click="editor.chain().focus().setColor('lime').run()" :class="{ 'is-active': editor.isActive('textStyle', { color: 'lime' })}"></button>
          <button class="color orange" @click="editor.chain().focus().setColor('orange').run()" :class="{ 'is-active': editor.isActive('textStyle', { color: 'orange' })}"></button>
        </div>
      </bubble-menu>
    </div>

    <editor-content :editor="editor"  @click="handleImagePlacement"/><br>
    <button style="border-radius: 8px;" @click="convertToPDF">Convert to PDF</button>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, onUnmounted } from 'vue';
import { Color } from '@tiptap/extension-color';
import Document from '@tiptap/extension-document';
import Paragraph from '@tiptap/extension-paragraph';
import Text from '@tiptap/extension-text';
import TextStyle from '@tiptap/extension-text-style';
import StarterKit from '@tiptap/starter-kit';
import TextAlign from '@tiptap/extension-text-align';
import Image from '@tiptap/extension-image';
import Underline from '@tiptap/extension-underline';
import { BubbleMenu, Editor, EditorContent, FloatingMenu } from '@tiptap/vue-3';

export default {
  components: {
    EditorContent,
    BubbleMenu,
    FloatingMenu,
  },
  
  setup() {
    const unsplashApiKey = import.meta.env.VITE_UNSPLASH_API_KEY;
    const editor = ref(null);
    const searchTerm = ref('');
    const images = ref([]);

    onMounted(() => {
      editor.value = new Editor({
        extensions: [
          StarterKit,
          Document,
          Paragraph,
          Text,
          TextStyle,
          Color,
          Image,
          TextAlign.configure({
            types: ['heading', 'paragraph', 'img'],
          }),
          Underline,
        ],
        content: `
          <h1>Hello world</h1>
        `,
      });
    });

    onUnmounted(() => {
      editor.value.destroy();
    });

    return {
      editor,
      searchTerm,
      images,
      unsplashApiKey,
    };
  },

  beforeUnmount() {
    this.editor.destroy()
  },
  methods: {
    handleImagePlacement() {
        const selection = this.editor.state.selection;
        if (selection && selection.$cursor) {
            const cursorPos = selection.$cursor.pos;
            const imageSrc = this.selectedImageSrc; 
        }
    },
    async searchImages() {
      try {
        const result = await this.SearchImages(this.searchTerm);
        this.images = result;
      } catch (error) {
        console.error('Error fetching images:', error);
      }
    },
    async convertToPDF() {
      try {
        const htmlContent = this.editor.getHTML();

        const requestBody = {
          htmlContent,
        };

        const response = await fetch('http://localhost:3000/service/convert', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(requestBody),
        });

        if (!response.ok) {
          throw new Error(`Failed to convert to PDF. Status: ${response.status}`);
        }

        const pdfBlob = await response.blob();
        const link = document.createElement('a');
        link.download = 'output.pdf';
        link.href = window.URL.createObjectURL(pdfBlob);
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
      } catch (error) {
        console.error('Error converting to PDF:', error);
      }
    },
    async SearchImages(search) {
      const api = await fetch(`https://api.unsplash.com/search/photos?page=1&query=${search}&client_id=${this.unsplashApiKey}`);
      const data = await api.json();
      return data.results;
    },
  },
  watch: {
    initialSearchTerm(newSearchTerm) {
      this.searchTerm = newSearchTerm;
    },
  },
  props: {
    initialSearchTerm: {
      type: String,
      default: '',
    },
  },
}


</script>

<style lang="scss">
/* Basic editor styles */
.tiptap {
  width:95%;
  border:1px solid white;
  background-color: white;

  > * + * {
    margin-top: 0.75em;
  }

  ul,
  ol {
    padding: 0 1rem;
  }

  blockquote {
    padding-left: 1rem;
    border-left: 2px solid rgba(#0D0D0D, 0.1);
  }
}

button{
  background-color:#ddd;
  padding:8px;
  border:none;
  color:#111;
}

.bubble-menu {
  display: flex;
  flex-wrap: wrap;
  background-color: #0D0D0D;
  padding: 0.2rem;
  border-radius: 0.5rem;

  .colors{
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    justify-content: center;
    gap:5px;
    width:100%;

      .color{
        display: block;
        width:calc(33.33% - 5px);
        height: 35px;
        border-radius: 0.5rem;
    }

  }

  .red{background-color:red}
  .blue{background-color: blue;}
  .green{background-color: green;}
  .yellow{background-color: yellow;}
  .orange{background-color: orange;}
  .lime{background-color: lime;}


  button {
    border: none;
    background: none;
    color: #FFF;
    font-size: 0.85rem;
    font-weight: 500;
    padding: 0 0.2rem;
    opacity: 0.8;
    width:33.33%;
    text-align: center;

    &:hover,
    &.is-active {
      opacity: 1;
    }
  }
}

.floating-menu {
  display: flex;
  background-color: #0D0D0D10;
  padding: 0.2rem;
  border-radius: 0.5rem;

  button {
    border: none;
    background: none;
    font-size: 0.85rem;
    font-weight: 500;
    padding: 0 0.2rem;
    opacity: 0.6;
    width:33.33%;

    &:hover,
    &.is-active {
      opacity: 1;
    }
  }
}
</style>