package com.ruoyi.traces.service.impl;

import com.ruoyi.system.mapper.SysUserMapper;
import com.ruoyi.traces.domain.TraceLngLat;
import com.ruoyi.traces.domain.TraceTransport;
import com.ruoyi.traces.domain.vo.LngLatVo;
import com.ruoyi.traces.mapper.TraceLngLatMapper;
import com.ruoyi.traces.mapper.TraceTransportMapper;
import com.ruoyi.traces.service.DriverService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class DriverServiceImpl implements DriverService {
    @Autowired
    private TraceTransportMapper traceTransportMapper;
    @Autowired
    private SysUserMapper sysUserMapper;
    @Autowired
    private TraceLngLatMapper traceLngLatMapper;
    @Override
    public List<TraceTransport> listTransport(String driverId) {
        return traceTransportMapper.listTransport(driverId);
    }

    @Override
    public int saveAndUpdateTransportInfoToDb(TraceLngLat traceLngLat) {
        TraceTransport traceTransport = new TraceTransport();
        traceTransport.setOutFactoryStatus(traceLngLat.getOutFactoryStatus());
        traceTransport.setId(traceLngLat.getId());
        traceTransport.setCropsId(traceLngLat.getCropsId());
        traceTransport.setStatus(traceLngLat.getStatus());

        traceLngLatMapper.insertSelective(traceLngLat);
        return traceTransportMapper.updateByCropsId(traceTransport);
    }

}
