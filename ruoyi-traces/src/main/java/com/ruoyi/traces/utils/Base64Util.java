package com.ruoyi.traces.utils;

import sun.misc.BASE64Decoder;

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.OutputStream;

public class Base64Util {
    public static File base64ConvertFile(String s) {
        String filePath = "/static";
        String fileName = System.currentTimeMillis()+".jpg";
        BASE64Decoder decoder = new BASE64Decoder();
        File file = null;
        try {
            byte[] bytes = decoder.decodeBuffer(s);
            for (int i = 0; i < bytes.length; ++i) {
                if (bytes[i] < 0) {
                    bytes[i] += 256;
                }
            }
            String imageFilePath = filePath+fileName.replace("\\\\","/");
            OutputStream out = new FileOutputStream(imageFilePath);
            out.write(bytes);
            out.flush();
            out.close();
            file = new File(imageFilePath);
        }catch (IOException e){
            e.printStackTrace();
        }
        return file;
    }
}
