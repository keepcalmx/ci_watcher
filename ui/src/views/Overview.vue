<script setup lang="ts">
import { ref } from 'vue'

const chartOptions = ref({
  chart: {
    id: "vuechart-example",
  },
  xaxis: {
    categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998],
  },
  title: {
    text: 'xx图分布'
  },
})
const series = ref(
  [
    {
      name: "series-1",
      data: [30, 40, 35, 50, 49, 60, 70, 91],
    },
  ]
)


// owner
const hot_map_series = ref([{
  name: '向峰',
  data: [0.10, 0.24, 0.36, 0.44, 0.52, 0.69, 0.77, 0.86, 0.70, 0.80, 0.82, 0.81, 0.81, 0.88, 0.92, 0.92, 0.92, 0.95, 0.98, 1]
},
{
  name: '唐佳乐',
  data: [0.20, 0.36, 0.57, 0.86, 0.29, 0.30, 0.75, 0.21, 0.87, 0.70, 0.52, 0.03, 0.16, 0.61, 0.98, 0.08, 0.59, 0.06, 0.69, 0.30]
},
{
  name: '王瑞琳',
  data: [0.17, 0.54, 0.54, 0.28, 0.42, 0.53, 0.25, 0.28, 0.79, 0.36, 0.88, 0.30, 0.89, 0.10, 0.98, 0.07, 0.22, 0.68, 0.24, 0.31]
},
{
  name: '许登辉',
  data: [0.93, 0.74, 0.80, 0.73, 0.18, 0.43, 0.90, 0.68, 0.98, 0.92, 0.09, 0.49, 0.93, 0.95, 0.35, 0.69, 0.71, 0.56, 0.65, 0.55]
},
{
  name: '魏俊明',
  data: [0.76, 0.40, 0.13, 0.99, 0.90, 0.32, 0.72, 0.64, 0.09, 0.67, 0.62, 0.37, 0.24, 0.54, 0.19, 0.24, 0.63, 0.13, 0.28, 0.41]
},
{
  name: '戴雪婷',
  data: [0.43, 0.63, 0.55, 0.62, 0.73, 0.83, 0.60, 0.74, 0.40, 0.50, 0.60, 0.41, 0.91, 0.90, 0.92, 0.92, 0.59, 0.56, 0.82, 0.88]
}
])

const hot_map_chartOptions = ref({
  xaxis: {
    categories: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20],
  },
  plotOptions: {
    heatmap: {
      // shadeIntensity: 0.5,
      radius: 0,
      useFillColorAsStroke: true,
    }
  },
  dataLabels: {
    enabled: true
  },
  colors: ["#008FFB"],
  stroke: {
    width: 1
  },
  title: {
    text: 'RECENT RESULT HEAT MAP'
  },
})


const type_bar_series = ref([{
  name: 'PASS',
  data: [22, 98, 77, 89, 50, 78]
}, {
  name: 'FAIL',
  data: [78, 2, 23, 11, 50, 22]
}])

const type_bar_chartOptions = ref({
  chart: {
    type: 'bar',
    height: 350,
    stacked: true,
    stackType: '100%'
  },
  plotOptions: {
    bar: {
      horizontal: true,
    },
  },
  stroke: {
    width: 1,
    colors: ['#fff']
  },
  title: {
    text: 'LAST RESULT By TYPE'
  },
  xaxis: {
    categories: ['CENT', 'CENT_SSH', 'CENT_SQL', 'DIST_SSH', 'DIST_SQL'],
  },
  fill: {
    opacity: 1
  },
  legend: {
    position: 'top',
    horizontalAlign: 'left',
    // offsetX: 40
  }
})


// owner
const bar_series = ref([{
  name: 'PASS',
  data: [100, 99, 100, 86, 50, 78]
}, {
  name: 'FAIL',
  data: [0, 1, 0, 14, 50, 22]
}])

const bar_chartOptions = ref({
  chart: {
    type: 'bar',
    height: 350,
    stacked: true,
    stackType: '100%'
  },
  plotOptions: {
    bar: {
      horizontal: true,
    },
  },
  stroke: {
    width: 1,
    colors: ['#fff']
  },
  title: {
    text: 'LAST RESULT By OWNER'
  },
  xaxis: {
    categories: ['向峰', '唐佳乐', '王瑞琳', '许登辉', '魏俊明', '戴雪婷'],
  },
  fill: {
    opacity: 1
  },
  legend: {
    position: 'top',
    horizontalAlign: 'left',
    // offsetX: 40
  }
})



const type_chartOptions = ref(
  {
    chart: {
      height: 390,
      type: 'radialBar',
    },
    plotOptions: {
      radialBar: {
        offsetY: 0,
        startAngle: 0,
        endAngle: 270,
        hollow: {
          margin: 5,
          size: '30%',
          background: 'transparent',
          image: undefined,
        },
        dataLabels: {
          name: {
            show: false,
          },
          value: {
            show: false,
          }
        }
      }
    },
    colors: ['#1ab7ea', '#0084ff', '#39539E', '#0077B5'],
    labels: ['Vimeo', 'Messenger', 'Facebook', 'LinkedIn'],
    legend: {
      show: true,
      floating: true,
      fontSize: '16px',
      position: 'left',
      offsetX: 160,
      offsetY: 15,
      labels: {
        useSeriesColors: true,
      },
      markers: {
        size: 0
      },
      formatter: function (seriesName: string, opts: { w: { globals: { series: { [x: string]: string; }; }; }; seriesIndex: string | number; }) {
        return seriesName + ":  " + opts.w.globals.series[opts.seriesIndex]
      },
      itemMargin: {
        vertical: 3
      }
    },
    responsive: [{
      breakpoint: 480,
      options: {
        legend: {
          show: false
        }
      }
    }]
  }
)

const cate_series = ref([44, 55, 13, 43, 22])
const cate_chartOptions = ref({
  chart: {
    // height: 400,
    // width: 380,
    type: 'pie',
  },
  labels: ['CENT', 'CENT_SSH', 'CENT_SQL', 'DIST_SSH', 'DIST_SQL'],
  responsive: [{
    breakpoint: 480,
    options: {
      chart: {
        width: 200
      },
      legend: {
        position: 'bottom'
      }
    }
  }]
})


const total_series = ref([87])
const total_chartOptions = ref({
  chart: {
    type: 'radialBar',
    offsetY: -10
  },
  plotOptions: {
    radialBar: {
      startAngle: -135,
      endAngle: 135,
      dataLabels: {
        name: {
          fontSize: '16px',
          color: undefined,
          offsetY: 120
        },
        value: {
          offsetY: 76,
          fontSize: '22px',
          color: undefined,
          formatter: function (val: string) {
            return val + "%";
          }
        }
      }
    }
  },
  fill: {
    type: 'gradient',
    gradient: {
      shade: 'dark',
      shadeIntensity: 0.15,
      inverseColors: false,
      opacityFrom: 1,
      opacityTo: 1,
      stops: [0, 50, 65, 91]
    },
  },
  stroke: {
    dashArray: 4
  },
  labels: ['TOTAL PASS'],
}
)
</script>

<template>
  <el-row justify="center" style="margin-top: 20px; margin-left: auto;">
    <el-col :span="6">
      <el-card style="max-height: 340px;">
        <apexchart type="radialBar" :options="total_chartOptions" :series="total_series"></apexchart>
      </el-card>
    </el-col>
    <el-col :span="6">
      <el-card style="min-height: 340px;">
        <!-- <span style="margin-top: 20px;"></span> -->
        <apexchart style="margin-top: 20px;" type="pie" height="600" :options="cate_chartOptions" :series="cate_series">
        </apexchart>
      </el-card>
    </el-col>
  </el-row>

  <el-row justify="center" style="margin-top: 20px;">
    <el-col :span="12">
      <el-card>
        <apexchart type="bar" height="350" :options="type_bar_chartOptions" :series="type_bar_series"></apexchart>
      </el-card>
    </el-col>
  </el-row>

  <el-row justify="center" style="margin-top: 20px;">
    <el-col :span="12">
      <el-card>
        <apexchart type="bar" height="350" :options="bar_chartOptions" :series="bar_series"></apexchart>
      </el-card>
    </el-col>
  </el-row>

  <el-row style="margin-top: 21px;"></el-row>

  <el-row justify="center">
    <el-col :span="12">
      <el-card>
        <apexchart type="heatmap" height="350" :options="hot_map_chartOptions" :series="hot_map_series"></apexchart>
      </el-card>
    </el-col>
  </el-row>
</template>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 25px;
}
</style>
