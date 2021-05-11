package com.ruoyi.traces.controller;

import com.ruoyi.common.core.domain.AjaxResult;
import com.ruoyi.traces.domain.TraceLngLat;
import com.ruoyi.traces.domain.vo.LngLatVo;
import com.ruoyi.traces.service.DriverService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/traces/driver")
public class DriverController {

    @Autowired
    private DriverService driverService;

    @GetMapping("/listTransport/{driverId}")
    public AjaxResult listTransport(@PathVariable("driverId") String driverId){
        return AjaxResult.success(driverService.listTransport(driverId));
    }

    @PostMapping("/saveAndUpdateTransportInfoToDb")
    public AjaxResult saveAndUpdateTransportInfoToDb(@RequestBody TraceLngLat traceLngLat){
        System.out.println("lat   "+traceLngLat);
        return AjaxResult.success(driverService.saveAndUpdateTransportInfoToDb(traceLngLat));
    }

}
