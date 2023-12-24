import { useMutation, useQueryClient } from "react-query";
import { docAPI } from "./axios";

export const useGenerateDoc = () => {
  const queryClient = useQueryClient();

  const generateDocMutation = useMutation(
    async (payload:string) => {
      const response = await docAPI().post("/doc", {docHTML:payload});
      return response.data; 
    },
    // {
    //   onSuccess: (data) => {
    //     const href = URL.createObjectURL(new Blob([data], { type: "application/pdf" }));
    //     const link = document.createElement("a");
    //     link.href = href;
    //     link.setAttribute("download", "meeting.pdf");
    //     document.body.appendChild(link);
    //     link.click();
    //     document.body.removeChild(link);
    //     URL.revokeObjectURL(href);
    //   },
    // }
  );

  return generateDocMutation;
};
