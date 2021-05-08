import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReportesMerkleComponent } from './reportes-merkle.component';

describe('ReportesMerkleComponent', () => {
  let component: ReportesMerkleComponent;
  let fixture: ComponentFixture<ReportesMerkleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReportesMerkleComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ReportesMerkleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
