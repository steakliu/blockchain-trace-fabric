package com.ruoyi.traces.controller;

import com.ruoyi.common.core.domain.AjaxResult;
import com.ruoyi.traces.domain.TraceTransport;
import com.ruoyi.traces.service.RetailerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/traces/retailer")
public class RetailerController {

    @Autowired
    private RetailerService retailerService;

    @PostMapping("/updateReceiveStatus")
    public AjaxResult updateReceiveStatus(@RequestBody TraceTransport traceTransport){
        return AjaxResult.success(retailerService.updateReceiveStatus(traceTransport));
    }
}
