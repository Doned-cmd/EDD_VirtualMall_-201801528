import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReportesRobotComponent } from './reportes-robot.component';

describe('ReportesRobotComponent', () => {
  let component: ReportesRobotComponent;
  let fixture: ComponentFixture<ReportesRobotComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReportesRobotComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ReportesRobotComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
