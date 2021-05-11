package com.ruoyi.traces.service;

import com.ruoyi.common.core.domain.AjaxResult;
import com.ruoyi.traces.domain.TraceTransport;

import java.util.List;

public interface MaterialService {
    List<TraceTransport> listCrops(String deptId);

    int changeInToOut(String cropsId);

    int addTransport(TraceTransport traceTransport);
}
