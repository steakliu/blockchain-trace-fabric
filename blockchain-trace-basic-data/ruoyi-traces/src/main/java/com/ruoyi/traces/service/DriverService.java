package com.ruoyi.traces.service;

import com.ruoyi.traces.domain.TraceLngLat;
import com.ruoyi.traces.domain.TraceTransport;
import com.ruoyi.traces.domain.vo.LngLatVo;

import java.util.List;

public interface DriverService {
    List<TraceTransport> listTransport(String driverId);

    int saveAndUpdateTransportInfoToDb(TraceLngLat traceLngLat);
}
