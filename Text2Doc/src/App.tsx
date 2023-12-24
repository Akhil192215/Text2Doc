import { useState } from 'react'

import './App.css'
import TextEditor from './component/Editor'



function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <div>
       <h3 className='text-red-700'>Hi</h3>
       <TextEditor/>
      </div>
     
    </>
  )
}

export default App
