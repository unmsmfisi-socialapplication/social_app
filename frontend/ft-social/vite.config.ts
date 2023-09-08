import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import CracoLessPuggin from 'craco-less'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [{
    ...react(),
    ...CracoLessPuggin,
  }],
})
