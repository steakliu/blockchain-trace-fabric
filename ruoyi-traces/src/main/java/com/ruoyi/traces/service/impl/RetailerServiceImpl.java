package com.ruoyi.traces.service.impl;

import com.ruoyi.traces.domain.TraceTransport;
import com.ruoyi.traces.mapper.TraceTransportMapper;
import com.ruoyi.traces.service.RetailerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class RetailerServiceImpl implements RetailerService {
    @Autowired
    private TraceTransportMapper traceTransportMapper;

    @Override
    public int updateReceiveStatus(TraceTransport traceTransport) {
        return traceTransportMapper.updateReceiveStatus(traceTransport);
    }
}
