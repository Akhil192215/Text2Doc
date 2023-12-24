import { useEffect, useState } from "react";

import "./App.css";
import { TextEditor } from "./component/Editor";
import { useGenerateDoc } from "./services/doc";


function App() {
  const [doc, setDoc] = useState("");
  const gnerateMutate = useGenerateDoc()

  useEffect(() => {
    console.log(doc);
  }, [doc]);

  const handlePDF = async() => {
    try {
     await gnerateMutate.mutateAsync(doc)
    } catch (error) {
      console.log(error);
      
    }
  };

  return (
    <>
      <div className="space-y-4">
        <h3 className="text-3xl font-semibold">Create custom document</h3>
       

        <TextEditor
          placeholder="Enter your document body"
          onChange={(e: any) => {
            setDoc(e);
          }}
        />
       <button onClick={handlePDF}>GeneratePDF</button>
      </div>
    </>
  );
}

export default App;
