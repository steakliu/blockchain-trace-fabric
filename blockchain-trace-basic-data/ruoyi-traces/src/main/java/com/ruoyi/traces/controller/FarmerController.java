package com.ruoyi.traces.controller;

import com.github.tobato.fastdfs.domain.fdfs.StorePath;
import com.github.tobato.fastdfs.service.FastFileStorageClient;
import com.ruoyi.common.annotation.RepeatSubmit;
import com.ruoyi.common.core.controller.BaseController;
import com.ruoyi.common.core.domain.AjaxResult;
import com.ruoyi.common.core.domain.entity.SysDept;
import com.ruoyi.common.core.domain.entity.SysUser;
import com.ruoyi.common.utils.SecurityUtils;
import com.ruoyi.traces.domain.TraceTransport;
import com.ruoyi.traces.domain.TraceUserCrops;
import com.ruoyi.traces.domain.vo.CropsImageVo;
import com.ruoyi.traces.service.FarmerService;
import com.ruoyi.traces.utils.Base64Util;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.*;
import sun.misc.BASE64Decoder;

import java.io.*;
import java.util.List;
import java.util.Map;

/**
 * 农户
 */
@RestController
@RequestMapping("/traces/farmer")
public class FarmerController extends BaseController {

    @Value("${fdfs.address}")
    private String address;

    @Autowired
    private FarmerService farmerService;
    @Autowired
    private FastFileStorageClient fastFileStorageClient;

    @PostMapping("/addTransport")
    public AjaxResult addTransport(@RequestBody TraceTransport traceTransport){
        return AjaxResult.success(farmerService.addTransport(traceTransport));
    }
    /**
     * @author wx
     * @Date  2020/11/3 0003
     * @Param [deptId]获取部门下的司机
     * @return com.ruoyi.common.core.domain.AjaxResult
     **/
    @GetMapping("/getAllDriverByDeptId/{deptId}")
    public AjaxResult getAllDriverByDeptId(@PathVariable("deptId") String deptId){
        List<SysUser> list = farmerService.getAllDriverByDeptId(deptId);
        return AjaxResult.success(list);
    }

    @GetMapping("/getFactoryByDeptId/{deptId}")
    public AjaxResult getFactoryByDeptId(@PathVariable("deptId") Integer deptId){
        return AjaxResult.success(farmerService.getFactoryByDeptId(deptId));
    }


    @PostMapping("imageUpload")
    public AjaxResult imageUpload(@RequestBody CropsImageVo cropsImageVo) throws FileNotFoundException {
        String imageBase64 = cropsImageVo.getImageBase64().replace("data:image/jpeg;base64,", "");
        File file = Base64Util.base64ConvertFile(imageBase64);
        StorePath storePath = this.fastFileStorageClient.uploadImageAndCrtThumbImage(new FileInputStream(file),file.length(),"jpg",null);
        return AjaxResult.success(address+storePath.getFullPath());
    }

    /**
     * 添加作物
     *
     * @param traceCrops
     * @return
     */
    @PostMapping("/saveCrops")
    @RepeatSubmit
    public AjaxResult saveCrops(@RequestBody TraceUserCrops traceCrops) {
        traceCrops.setUsername(SecurityUtils.getUsername());
        return AjaxResult.success(farmerService.saveCrops(traceCrops));
    }

    /**
     * @return com.ruoyi.common.core.domain.AjaxResult
     * @author steakliu
     * @Date 23:12 2020/10/27 0027
     * @Param [username]
     **/
    @GetMapping("/getCropsByUsername/{username}")
    public AjaxResult getCropsByUsername(@PathVariable("username") String username) {
        return AjaxResult.success(farmerService.getCropsByUsername(username));
    }

}
