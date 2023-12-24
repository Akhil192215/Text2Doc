import { useEffect, useState } from "react";

import "./App.css";
import TextEditor from "./component/Editor";

function App() {
  const [doc, setDoc] = useState("");

  useEffect(() => {
    console.log(doc);
  }, [doc]);

  const handleChange = (e) => {
    console.log(e);
  };

  return (
    <>
      <div className="space-y-4">
        <h3 className="text-3xl font-semibold">Create custom document</h3>
        <TextEditor
          placeholder="Enter your doc body"
          onChange={(e: any) => {
            setDoc(e);
          }}
        />
      </div>
    </>
  );
}

export default App;
