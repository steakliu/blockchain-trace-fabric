package com.ruoyi.traces.controller;

import com.ruoyi.common.core.domain.AjaxResult;
import com.ruoyi.traces.domain.TraceTransport;
import com.ruoyi.traces.service.MaterialService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/traces/material")
public class MaterialController {
    @Autowired
    private MaterialService materialService;

    @GetMapping("/listCrops/{deptId}")
    public AjaxResult listCrops(@PathVariable("deptId") String deptId){
        return AjaxResult.success(materialService.listCrops(deptId));
    }

    /**
     * @author wx
     * @Date 23:47 2020/11/7 0007
     * @Param [cropsId] 出库
     * @return com.ruoyi.common.core.domain.AjaxResult
     **/
    @GetMapping("/changeInToOut/{cropsId}")
    public AjaxResult changeInToOut(@PathVariable("cropsId") String cropsId){
        return AjaxResult.success(materialService.changeInToOut(cropsId));
    }

    @PostMapping("/addTransport")
    public AjaxResult addTransport(@RequestBody TraceTransport traceTransport){
        return AjaxResult.success(materialService.addTransport(traceTransport));
    }
}
